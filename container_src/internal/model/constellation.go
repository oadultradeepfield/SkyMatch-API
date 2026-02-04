package model

import "strings"

const (
	noirLabBaseURL = "https://storage.noirlab.edu/media/archives/images/thumb350x/"
	noirLabSuffix  = "-ann.jpg"
)

type Constellation struct {
	LatinName   string
	EnglishName string
	ImageID     string
}

func (c Constellation) ImageURL() string {
	return noirLabBaseURL + c.ImageID + noirLabSuffix
}

var Constellations = []Constellation{
	{LatinName: "Andromeda", EnglishName: "Princess of Ethiopia", ImageID: "andromeda"},
	{LatinName: "Antlia", EnglishName: "Air Pump", ImageID: "antlia"},
	{LatinName: "Apus", EnglishName: "Bird of Paradise", ImageID: "apus"},
	{LatinName: "Aquarius", EnglishName: "Water Bearer", ImageID: "aquarius"},
	{LatinName: "Aquila", EnglishName: "Eagle", ImageID: "aquila"},
	{LatinName: "Ara", EnglishName: "Altar", ImageID: "ara"},
	{LatinName: "Aries", EnglishName: "Ram", ImageID: "aries"},
	{LatinName: "Auriga", EnglishName: "Charioteer", ImageID: "auriga"},
	{LatinName: "Boötes", EnglishName: "Herdsman", ImageID: "bootes"},
	{LatinName: "Caelum", EnglishName: "Chisel", ImageID: "caelum"},
	{LatinName: "Camelopardalis", EnglishName: "Giraffe", ImageID: "camelopardalis"},
	{LatinName: "Cancer", EnglishName: "Crab", ImageID: "cancer"},
	{LatinName: "Canes Venatici", EnglishName: "Hunting Dogs", ImageID: "canesVenatici"},
	{LatinName: "Canis Major", EnglishName: "Great Dog", ImageID: "canis-major"},
	{LatinName: "Canis Minor", EnglishName: "Little Dog", ImageID: "canis-minor"},
	{LatinName: "Capricornus", EnglishName: "Sea Goat", ImageID: "capricornus"},
	{LatinName: "Carina", EnglishName: "Keel", ImageID: "carina"},
	{LatinName: "Cassiopeia", EnglishName: "Queen of Ethiopia", ImageID: "cassiopeia"},
	{LatinName: "Centaurus", EnglishName: "Centaur", ImageID: "centaurus"},
	{LatinName: "Cepheus", EnglishName: "King of Ethiopia", ImageID: "cepheus"},
	{LatinName: "Cetus", EnglishName: "Whale", ImageID: "cetus"},
	{LatinName: "Chamaeleon", EnglishName: "Chameleon", ImageID: "chamaeleon"},
	{LatinName: "Circinus", EnglishName: "Compass", ImageID: "circinus"},
	{LatinName: "Columba", EnglishName: "Dove", ImageID: "columba"},
	{LatinName: "Coma Berenices", EnglishName: "Berenice's Hair", ImageID: "coma-berenices"},
	{LatinName: "Corona Australis", EnglishName: "Southern Crown", ImageID: "corona-australis"},
	{LatinName: "Corona Borealis", EnglishName: "Northern Crown", ImageID: "coronaBorealis"},
	{LatinName: "Corvus", EnglishName: "Crow", ImageID: "corvus"},
	{LatinName: "Crater", EnglishName: "Cup", ImageID: "crater"},
	{LatinName: "Crux", EnglishName: "Southern Cross", ImageID: "crux"},
	{LatinName: "Cygnus", EnglishName: "Swan", ImageID: "cygnus"},
	{LatinName: "Delphinus", EnglishName: "Dolphin", ImageID: "delphinus"},
	{LatinName: "Dorado", EnglishName: "Swordfish", ImageID: "dorado"},
	{LatinName: "Draco", EnglishName: "Dragon", ImageID: "draco"},
	{LatinName: "Equuleus", EnglishName: "Little Horse", ImageID: "equuleus"},
	{LatinName: "Eridanus", EnglishName: "River", ImageID: "eridanus"},
	{LatinName: "Fornax", EnglishName: "Furnace", ImageID: "fornax"},
	{LatinName: "Gemini", EnglishName: "Twins", ImageID: "gemini"},
	{LatinName: "Grus", EnglishName: "Crane", ImageID: "grus"},
	{LatinName: "Hercules", EnglishName: "Hercules", ImageID: "hercules"},
	{LatinName: "Horologium", EnglishName: "Clock", ImageID: "horologium"},
	{LatinName: "Hydra", EnglishName: "Sea Serpent", ImageID: "hydra"},
	{LatinName: "Hydrus", EnglishName: "Water Snake", ImageID: "hydrus"},
	{LatinName: "Indus", EnglishName: "Indian", ImageID: "indus"},
	{LatinName: "Lacerta", EnglishName: "Lizard", ImageID: "lacerta"},
	{LatinName: "Leo", EnglishName: "Lion", ImageID: "leo"},
	{LatinName: "Leo Minor", EnglishName: "Little Lion", ImageID: "leo-minor"},
	{LatinName: "Lepus", EnglishName: "Hare", ImageID: "lepus"},
	{LatinName: "Libra", EnglishName: "Scales", ImageID: "libra"},
	{LatinName: "Lupus", EnglishName: "Wolf", ImageID: "lupus"},
	{LatinName: "Lynx", EnglishName: "Lynx", ImageID: "lynx"},
	{LatinName: "Lyra", EnglishName: "Lyre", ImageID: "lyra"},
	{LatinName: "Mensa", EnglishName: "Table Mountain", ImageID: "mensa"},
	{LatinName: "Microscopium", EnglishName: "Microscope", ImageID: "microscopium"},
	{LatinName: "Monoceros", EnglishName: "Unicorn", ImageID: "monoceros"},
	{LatinName: "Musca", EnglishName: "Fly", ImageID: "musca"},
	{LatinName: "Norma", EnglishName: "Carpenter's Square", ImageID: "norma"},
	{LatinName: "Octans", EnglishName: "Octant", ImageID: "octans"},
	{LatinName: "Ophiuchus", EnglishName: "Serpent Bearer", ImageID: "ophiuchus"},
	{LatinName: "Orion", EnglishName: "Hunter", ImageID: "orion"},
	{LatinName: "Pavo", EnglishName: "Peacock", ImageID: "pavo"},
	{LatinName: "Pegasus", EnglishName: "Winged Horse", ImageID: "pegasus"},
	{LatinName: "Perseus", EnglishName: "Hero", ImageID: "perseus"},
	{LatinName: "Phoenix", EnglishName: "Phoenix", ImageID: "phoenix"},
	{LatinName: "Pictor", EnglishName: "Painter's Easel", ImageID: "pictor"},
	{LatinName: "Pisces", EnglishName: "Fish", ImageID: "pisces"},
	{LatinName: "Piscis Austrinus", EnglishName: "Southern Fish", ImageID: "piscis-austrinus"},
	{LatinName: "Puppis", EnglishName: "Stern", ImageID: "puppis"},
	{LatinName: "Pyxis", EnglishName: "Compass", ImageID: "pyxis"},
	{LatinName: "Reticulum", EnglishName: "Net", ImageID: "reticulum"},
	{LatinName: "Sagitta", EnglishName: "Arrow", ImageID: "sagitta"},
	{LatinName: "Sagittarius", EnglishName: "Archer", ImageID: "sagittarius"},
	{LatinName: "Scorpius", EnglishName: "Scorpion", ImageID: "scorpius"},
	{LatinName: "Sculptor", EnglishName: "Sculptor", ImageID: "sculptor"},
	{LatinName: "Scutum", EnglishName: "Shield", ImageID: "scutum"},
	{LatinName: "Serpens", EnglishName: "Serpent", ImageID: "serpens-cauda"},
	{LatinName: "Sextans", EnglishName: "Sextant", ImageID: "sextans"},
	{LatinName: "Taurus", EnglishName: "Bull", ImageID: "taurus"},
	{LatinName: "Telescopium", EnglishName: "Telescope", ImageID: "telescopium"},
	{LatinName: "Triangulum", EnglishName: "Triangle", ImageID: "triangulum"},
	{LatinName: "Triangulum Australe", EnglishName: "Southern Triangle", ImageID: "triangulum-australe"},
	{LatinName: "Tucana", EnglishName: "Toucan", ImageID: "tucana"},
	{LatinName: "Ursa Major", EnglishName: "Great Bear", ImageID: "ursa-major"},
	{LatinName: "Ursa Minor", EnglishName: "Little Bear", ImageID: "ursa-minor"},
	{LatinName: "Vela", EnglishName: "Sails", ImageID: "vela"},
	{LatinName: "Virgo", EnglishName: "Virgin", ImageID: "virgo"},
	{LatinName: "Volans", EnglishName: "Flying Fish", ImageID: "volans"},
	{LatinName: "Vulpecula", EnglishName: "Fox", ImageID: "vulpecula"},
}

// SearchConstellations filters constellations by query (case-insensitive)
func SearchConstellations(query string) []Constellation {
	if query == "" {
		return Constellations
	}
	var results []Constellation
	for _, c := range Constellations {
		if strings.Contains(strings.ToLower(c.LatinName), query) ||
			strings.Contains(strings.ToLower(c.EnglishName), query) {
			results = append(results, c)
		}
	}
	return results
}
