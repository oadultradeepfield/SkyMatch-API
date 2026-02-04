package solve

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"server/internal/client/nova"
	"server/internal/client/simbad"
	"server/internal/model"
)

type Service struct {
	nova   *nova.Client
	simbad *simbad.Client
}

func NewService() *Service {
	return &Service{nova: nova.NewClient(), simbad: simbad.NewClient()}
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
			annMap[name] = a
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
	obj := &model.IdentifiedObject{Identifier: name, Name: name}

	if ann, ok := annMap[name]; ok {
		obj.XCoordinate = ann.PixelX
		obj.YCoordinate = ann.PixelY
	}

	info, err := s.simbad.QueryObject(name)
	if err != nil {
		obj.Type = classifyByName(name)
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

func classifyByName(name string) model.ObjectType {
	lower := strings.ToLower(name)
	for _, p := range []string{"m ", "ngc ", "ic "} {
		if strings.HasPrefix(lower, p) {
			return model.ObjectTypeDSO
		}
	}
	return model.ObjectTypeStar
}

func classifyByType(t string) model.ObjectType {
	lower := strings.ToLower(t)
	for _, k := range []string{"galaxy", "nebula", "cluster", "hii", "supernova"} {
		if strings.Contains(lower, k) {
			return model.ObjectTypeDSO
		}
	}
	return model.ObjectTypeStar
}

func classifyDSOType(t string) model.DeepSkyObjectType {
	lower := strings.ToLower(t)
	switch {
	case strings.Contains(lower, "galaxy"):
		return model.DSOGalaxy
	case strings.Contains(lower, "open") && strings.Contains(lower, "cluster"):
		return model.DSOOpenCluster
	case strings.Contains(lower, "globular"):
		return model.DSOGlobularCluster
	case strings.Contains(lower, "supernova"):
		return model.DSOSupernova
	default:
		return model.DSONebula
	}
}

func parseSpectralClass(sp string) model.SpectralClass {
	if len(sp) == 0 {
		return ""
	}
	switch sp[0] {
	case 'O':
		return model.SpectralO
	case 'B':
		return model.SpectralB
	case 'A':
		return model.SpectralA
	case 'F':
		return model.SpectralF
	case 'G':
		return model.SpectralG
	case 'K':
		return model.SpectralK
	case 'M':
		return model.SpectralM
	default:
		return ""
	}
}
