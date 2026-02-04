package solve

import (
	"strings"

	"server/internal/client/nova"
)

// Nova API returns object names in inconsistent formats that need cleaning:
// - Malformed names like "34 Ori)" or "Betelgeux (α Ori" (unbalanced parens)
// - Prefixes like "The star Betelgeuse"
// - Non-object entries like "Part of the constellation Orion (Ori)"
// - Annotation names use slash separators: "α Ori / 58 Ori", "Betelgeuse / Al Mankib / Betelgeux"
// - objects_in_field uses different formats: "Betelgeuse", "σ Ori", "Alnilam (ε Ori"

func cleanObjectName(name string) string {
	name = strings.TrimPrefix(name, "The star ")

	opening := strings.Count(name, "(")
	closing := strings.Count(name, ")")
	if closing > opening {
		name = strings.TrimSuffix(name, ")")
	} else if opening > closing {
		name = name + ")"
	}
	return name
}

func splitAnnotationNames(name string) []string {
	return strings.Split(name, " / ")
}

func extractNameWithoutParen(name string) string {
	if idx := strings.Index(name, " ("); idx != -1 {
		return name[:idx]
	}
	return name
}

func extractParenContent(name string) string {
	start := strings.Index(name, "(")
	end := strings.Index(name, ")")
	if start != -1 && end > start {
		return name[start+1 : end]
	}
	return ""
}

func shouldSkipObject(name string) bool {
	return strings.HasPrefix(name, "Part of the constellation")
}

func lookupAnnotation(name string, annMap map[string]nova.Annotation) (nova.Annotation, bool) {
	if ann, ok := annMap[name]; ok {
		return ann, true
	}
	if base := extractNameWithoutParen(name); base != name {
		if ann, ok := annMap[base]; ok {
			return ann, true
		}
	}
	if paren := extractParenContent(name); paren != "" {
		if ann, ok := annMap[paren]; ok {
			return ann, true
		}
	}
	return nova.Annotation{}, false
}

var greekToLatin = map[string]string{
	"α": "alf", "β": "bet", "γ": "gam", "δ": "del", "ε": "eps", "ζ": "zet",
	"η": "eta", "θ": "tet", "ι": "iot", "κ": "kap", "λ": "lam", "μ": "mu.",
	"ν": "nu.", "ξ": "ksi", "ο": "omi", "π": "pi.", "ρ": "rho", "σ": "sig",
	"τ": "tau", "υ": "ups", "φ": "phi", "χ": "chi", "ψ": "psi", "ω": "ome",
}

func greekToSimbadName(name string) string {
	for greek, latin := range greekToLatin {
		if strings.Contains(name, greek) {
			return strings.Replace(name, greek, "* "+latin, 1)
		}
	}
	return ""
}
