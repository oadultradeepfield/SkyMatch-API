package solve

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"

	"server/internal/client"
	"server/internal/client/nova"
	apperrors "server/internal/errors"
	"server/internal/model"
	"server/internal/service"
)

var _ service.SolveService = (*Service)(nil)

type Service struct {
	nova   client.NovaClient
	simbad client.SimbadClient
	apiKey string
}

func NewService(novaClient client.NovaClient, simbadClient client.SimbadClient, apiKey string) *Service {
	return &Service{nova: novaClient, simbad: simbadClient, apiKey: apiKey}
}

func (s *Service) Submit(ctx context.Context, file io.Reader, filename string) (int, error) {
	if s.apiKey == "" {
		return 0, apperrors.NewValidationError("NOVA_API_KEY not set")
	}
	session, err := s.nova.Login(ctx, s.apiKey)
	if err != nil {
		return 0, apperrors.NewExternalError("nova", err)
	}
	subID, err := s.nova.Upload(ctx, session, file, filename)
	if err != nil {
		return 0, apperrors.NewExternalError("nova", err)
	}
	return subID, nil
}

func (s *Service) GetStatus(ctx context.Context, subID int, fetch bool) (*model.SolveResult, error) {
	result := &model.SolveResult{JobID: fmt.Sprintf("%d", subID)}

	sub, err := s.nova.GetSubmission(ctx, subID)
	if err != nil {
		result.Status = model.StatusFailure
		return result, nil
	}

	if len(sub.Jobs) == 0 {
		result.Status = model.StatusQueued
		return result, nil
	}

	jobID := sub.Jobs[0]
	result.NovaJobID = jobID

	status, err := s.nova.GetJobStatus(ctx, jobID)
	if err != nil {
		result.Status = model.StatusFailure
		return result, nil
	}

	switch status {
	case "failure":
		result.Status = model.StatusFailure
		return result, nil
	case "success":
		if !fetch {
			result.Status = model.StatusGettingMoreDetails
			return result, nil
		}
	default:
		result.Status = model.StatusIdentifyingObjects
		return result, nil
	}

	return s.fetchFullData(ctx, subID, jobID)
}

func (s *Service) fetchFullData(ctx context.Context, subID, jobID int) (*model.SolveResult, error) {
	result := &model.SolveResult{
		JobID:     fmt.Sprintf("%d", subID),
		NovaJobID: jobID,
	}

	info, err := s.nova.GetJobInfo(ctx, jobID)
	if err != nil || len(info.ObjectsInField) == 0 {
		result.Status = model.StatusFailure
		return result, nil
	}

	annotations, err := s.nova.GetAnnotations(ctx, jobID)
	if err != nil {
		result.Status = model.StatusFailure
		return result, nil
	}

	annMap := make(map[string]nova.Annotation)
	for _, a := range annotations {
		for _, name := range a.Names {
			for _, part := range splitAnnotationNames(name) {
				annMap[part] = a
			}
		}
	}

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(10)
	var mu sync.Mutex

	for _, name := range info.ObjectsInField {
		name := name
		g.Go(func() error {
			obj, err := s.processObject(ctx, name, annMap)
			if err != nil {
				log.Printf("process %s: %v", name, err)
				return nil
			}
			if obj != nil {
				mu.Lock()
				result.Objects = append(result.Objects, *obj)
				mu.Unlock()
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	result.Status = model.StatusSuccess
	result.AnnotatedImageURL = s.nova.AnnotatedImageURL(jobID)
	return result, nil
}

func (s *Service) processObject(ctx context.Context, name string, annMap map[string]nova.Annotation) (*model.IdentifiedObject, error) {
	if shouldSkipObject(name) {
		return nil, nil
	}

	cleanedName := cleanObjectName(name)
	obj := &model.IdentifiedObject{Identifier: cleanedName, Name: cleanedName}

	if ann, ok := lookupAnnotation(cleanedName, annMap); ok {
		obj.XCoordinate = ann.PixelX
		obj.YCoordinate = ann.PixelY
	}

	info, err := s.simbad.QueryObject(ctx, cleanedName)
	if err != nil {
		if base := extractNameWithoutParen(cleanedName); base != cleanedName {
			info, err = s.simbad.QueryObject(ctx, base)
		}
	}
	if err != nil {
		if simbadName := greekToSimbadName(cleanedName); simbadName != "" {
			info, err = s.simbad.QueryObject(ctx, simbadName)
		}
	}
	if err != nil {
		obj.Type = classifyByName(cleanedName)
		return obj, nil
	}

	obj.Type = classifyByType(info.ObjectType)

	if obj.Type == model.ObjectTypeStar {
		if info.VMagnitude != nil && *info.VMagnitude < 3.0 {
			obj.VMagnitude = info.VMagnitude
			obj.SpectralClass = parseSpectralClass(info.SpectralType)
			obj.DistanceParsecs = info.DistanceParsecs()
		}
	} else {
		obj.DSOType = classifyDSOType(info.ObjectType)
	}

	if info.RA != nil && info.Dec != nil {
		obj.Constellation = model.GetConstellationByCoords(*info.RA, *info.Dec)
	}

	return obj, nil
}
