package solve

import (
	"fmt"
	"io"
	"os"
	"sync"

	"server/internal/client"
	"server/internal/client/nova"
	"server/internal/model"
)

type Service struct {
	nova   client.NovaClient
	simbad client.SimbadClient
}

func NewService(novaClient client.NovaClient, simbadClient client.SimbadClient) *Service {
	return &Service{nova: novaClient, simbad: simbadClient}
}

func (s *Service) Submit(file io.Reader, filename string) (int, error) {
	apiKey := os.Getenv("NOVA_API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("NOVA_API_KEY not set")
	}
	session, err := s.nova.Login(apiKey)
	if err != nil {
		return 0, err
	}
	return s.nova.Upload(session, file, filename)
}

func (s *Service) GetStatus(subID int, fetch bool) (*model.SolveResult, error) {
	result := &model.SolveResult{JobID: fmt.Sprintf("%d", subID)}

	sub, err := s.nova.GetSubmission(subID)
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

	status, err := s.nova.GetJobStatus(jobID)
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

	return s.fetchFullData(subID, jobID)
}

func (s *Service) fetchFullData(subID, jobID int) (*model.SolveResult, error) {
	result := &model.SolveResult{
		JobID:     fmt.Sprintf("%d", subID),
		NovaJobID: jobID,
	}

	info, err := s.nova.GetJobInfo(jobID)
	if err != nil || len(info.ObjectsInField) == 0 {
		result.Status = model.StatusFailure
		return result, nil
	}

	annotations, err := s.nova.GetAnnotations(jobID)
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

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, name := range info.ObjectsInField {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			if obj := s.processObject(n, annMap); obj != nil {
				mu.Lock()
				result.Objects = append(result.Objects, *obj)
				mu.Unlock()
			}
		}(name)
	}

	wg.Wait()

	result.Status = model.StatusSuccess
	result.AnnotatedImageURL = s.nova.AnnotatedImageURL(jobID)
	return result, nil
}

func (s *Service) processObject(name string, annMap map[string]nova.Annotation) *model.IdentifiedObject {
	if shouldSkipObject(name) {
		return nil
	}

	cleanedName := cleanObjectName(name)
	obj := &model.IdentifiedObject{Identifier: cleanedName, Name: cleanedName}

	if ann, ok := lookupAnnotation(cleanedName, annMap); ok {
		obj.XCoordinate = ann.PixelX
		obj.YCoordinate = ann.PixelY
	}

	info, err := s.simbad.QueryObject(cleanedName)
	if err != nil {
		if base := extractNameWithoutParen(cleanedName); base != cleanedName {
			info, err = s.simbad.QueryObject(base)
		}
	}
	if err != nil {
		if simbadName := greekToSimbadName(cleanedName); simbadName != "" {
			info, err = s.simbad.QueryObject(simbadName)
		}
	}
	if err != nil {
		obj.Type = classifyByName(cleanedName)
		return obj
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

	return obj
}
