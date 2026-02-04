package solve

import (
	"strings"

	"server/internal/model"
)

func classifyByName(name string) model.ObjectType {
	lower := strings.ToLower(name)
	for _, p := range []string{"m ", "ngc ", "ic "} {
		if strings.HasPrefix(lower, p) {
			return model.ObjectTypeDSO
		}
	}
	if strings.HasSuffix(lower, "nebula") {
		return model.ObjectTypeDSO
	}
	return model.ObjectTypeStar
}

func classifyByType(t string) model.ObjectType {
	lower := strings.ToLower(t)
	for _, k := range []string{"galaxy", "nebula", "cluster", "hii", "supernova", "cl*", "g ", "gxy", "snr"} {
		if strings.Contains(lower, k) {
			return model.ObjectTypeDSO
		}
	}
	return model.ObjectTypeStar
}

func classifyDSOType(t string) model.DeepSkyObjectType {
	lower := strings.ToLower(t)
	switch {
	case strings.Contains(lower, "galaxy") || strings.Contains(lower, "gxy") || lower == "g":
		return model.DSOGalaxy
	case strings.Contains(lower, "open") && strings.Contains(lower, "cluster"):
		return model.DSOOpenCluster
	case strings.Contains(lower, "globular") || lower == "glc":
		return model.DSOGlobularCluster
	case strings.Contains(lower, "supernova") || strings.Contains(lower, "snr"):
		return model.DSOSupernova
	case strings.Contains(lower, "cl*") || strings.Contains(lower, "cluster"):
		return model.DSOOpenCluster
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
