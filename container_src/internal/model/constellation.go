package model

import "strings"

const (
	noirLabBaseURL = "https://storage.noirlab.edu/media/archives/images/thumb350x/"
	noirLabSuffix  = "-ann.jpg"
)

type Constellation struct {
	Abbr        string
	LatinName   string
	EnglishName string
	ImageID     string
	Boundaries  []Boundary
}

type Boundary struct {
	RALow, RAHigh, DecLow float64
}

func (c *Constellation) ImageURL() string {
	return noirLabBaseURL + c.ImageID + noirLabSuffix
}

func (c *Constellation) Contains(ra, dec float64) bool {
	for ra < 0 {
		ra += 360
	}
	for ra >= 360 {
		ra -= 360
	}
	raHours := ra / 15.0

	for _, b := range c.Boundaries {
		if dec >= b.DecLow && raHours >= b.RALow && raHours < b.RAHigh {
			return true
		}
	}
	return false
}

func GetConstellationByCoords(ra, dec float64) *Constellation {
	for i := range Constellations {
		if Constellations[i].Contains(ra, dec) {
			return &Constellations[i]
		}
	}
	return nil
}

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

var Constellations = []Constellation{
	{Abbr: "And", LatinName: "Andromeda", EnglishName: "Princess of Ethiopia", ImageID: "andromeda", Boundaries: []Boundary{{23.75, 24.0, 24.0}, {0.0, 0.1, 23.5}, {0.1, 0.85, 23.0}, {0.85, 0.983, 20.0}, {0.983, 1.75, 16.167}}},
	{Abbr: "Ant", LatinName: "Antlia", EnglishName: "Air Pump", ImageID: "antlia", Boundaries: []Boundary{{9.333, 9.917, -39.75}}},
	{Abbr: "Aps", LatinName: "Apus", EnglishName: "Bird of Paradise", ImageID: "apus", Boundaries: []Boundary{{15.333, 17.5, -80.0}, {14.167, 15.333, -76.0}, {16.5, 17.5, -74.5}, {9.333, 17.5, -90.0}}},
	{Abbr: "Aqr", LatinName: "Aquarius", EnglishName: "Water Bearer", ImageID: "aquarius", Boundaries: []Boundary{{20.833, 21.333, -9.0}, {21.333, 21.867, -10.0}, {21.867, 23.0, -25.5}, {20.333, 21.333, -24.5}, {21.333, 21.867, -24.5}, {22.75, 23.833, -20.0}, {21.867, 22.75, -18.5}}},
	{Abbr: "Aql", LatinName: "Aquila", EnglishName: "Eagle", ImageID: "aquila", Boundaries: []Boundary{{18.867, 19.25, 6.25}, {19.167, 19.5, -3.5}, {19.5, 20.0, -4.0}, {20.0, 20.5, -5.0}}},
	{Abbr: "Ara", LatinName: "Ara", EnglishName: "Altar", ImageID: "ara", Boundaries: []Boundary{{16.625, 17.0, -50.0}, {17.0, 17.5, -54.0}, {17.5, 18.0, -55.333}, {18.0, 18.5, -56.0}, {18.5, 19.0, -57.0}, {16.625, 17.0, -58.0}, {17.0, 17.5, -58.5}, {17.5, 18.0, -59.0}, {16.5, 17.0, -67.5}, {17.0, 17.5, -68.0}, {17.5, 18.0, -70.0}, {16.5, 17.5, -74.5}, {16.5, 17.5, -77.0}}},
	{Abbr: "Ari", LatinName: "Aries", EnglishName: "Ram", ImageID: "aries", Boundaries: []Boundary{{1.917, 2.0, 27.25}, {2.0, 2.5, 26.5}, {2.5, 2.717, 22.833}}},
	{Abbr: "Aur", LatinName: "Auriga", EnglishName: "Charioteer", ImageID: "auriga", Boundaries: []Boundary{{5.7, 5.883, 32.0}, {5.167, 5.7, 28.0}, {5.7, 5.883, 26.5}}},
	{Abbr: "Boo", LatinName: "Bootes", EnglishName: "Herdsman", ImageID: "bootes", Boundaries: []Boundary{{14.0, 15.25, 42.0}, {14.0, 14.667, 36.0}, {14.667, 15.0, 24.0}, {15.0, 15.333, 23.5}}},
	{Abbr: "Cae", LatinName: "Caelum", EnglishName: "Chisel", ImageID: "caelum", Boundaries: []Boundary{{5.5, 6.0, -44.0}, {6.0, 6.5, -44.5}}},
	{Abbr: "Cam", LatinName: "Camelopardalis", EnglishName: "Giraffe", ImageID: "camelopardalis", Boundaries: []Boundary{{9.167, 10.667, 82.0}, {11.5, 13.583, 77.0}, {7.967, 9.167, 73.5}, {3.1, 3.167, 57.0}, {3.167, 3.333, 55.0}, {3.333, 5.0, 52.5}, {5.0, 6.1, 56.0}, {6.1, 7.0, 62.0}, {7.0, 7.967, 60.0}}},
	{Abbr: "Cnc", LatinName: "Cancer", EnglishName: "Crab", ImageID: "cancer", Boundaries: []Boundary{{8.083, 8.4, 23.0}, {8.4, 8.667, 21.0}, {8.667, 9.0, 18.5}, {9.0, 9.25, 16.167}}},
	{Abbr: "CVn", LatinName: "Canes Venatici", EnglishName: "Hunting Dogs", ImageID: "canesVenatici", Boundaries: []Boundary{{12.0, 14.0, 45.5}}},
	{Abbr: "CMa", LatinName: "Canis Major", EnglishName: "Great Dog", ImageID: "canis-major", Boundaries: []Boundary{{6.583, 7.667, -40.0}}},
	{Abbr: "CMi", LatinName: "Canis Minor", EnglishName: "Little Dog", ImageID: "canis-minor", Boundaries: []Boundary{{7.25, 7.5, 12.5}, {7.0, 7.25, 12.5}}},
	{Abbr: "Cap", LatinName: "Capricornus", EnglishName: "Sea Goat", ImageID: "capricornus", Boundaries: []Boundary{{20.0, 20.333, -22.0}, {20.333, 21.333, -24.5}, {21.333, 21.867, -28.833}, {20.333, 21.333, -33.0}}},
	{Abbr: "Car", LatinName: "Carina", EnglishName: "Keel", ImageID: "carina", Boundaries: []Boundary{{7.5, 7.833, -55.5}, {7.5, 7.833, -57.5}, {7.5, 7.833, -60.0}, {7.167, 7.5, -63.5}, {6.0, 6.5, -64.0}, {6.5, 6.833, -64.0}, {7.5, 7.833, -70.0}, {7.667, 8.333, -71.5}, {8.333, 9.0, -72.5}, {9.0, 9.333, -73.5}, {9.333, 10.583, -74.167}, {10.583, 11.833, -80.5}, {9.333, 10.583, -82.5}, {10.583, 11.833, -84.5}, {9.333, 10.583, -85.5}, {10.583, 11.833, -86.5}, {9.333, 11.833, -88.5}}},
	{Abbr: "Cas", LatinName: "Cassiopeia", EnglishName: "Queen of Ethiopia", ImageID: "cassiopeia", Boundaries: []Boundary{{3.1, 3.417, 68.0}, {0.0, 2.433, 58.5}, {1.7, 1.9, 57.5}, {2.433, 3.1, 57.0}, {22.867, 23.583, 59.083}, {22.317, 22.867, 56.25}, {22.133, 22.317, 55.0}, {22.867, 23.333, 52.5}, {23.333, 24.0, 50.0}, {0.0, 0.167, 48.0}, {23.583, 24.0, 48.0}, {0.167, 0.867, 46.0}, {0.0, 0.167, 48.0}, {0.167, 0.867, 50.5}, {0.867, 1.1, 50.5}}},
	{Abbr: "Cen", LatinName: "Centaurus", EnglishName: "Centaur", ImageID: "centaurus", Boundaries: []Boundary{{13.5, 14.167, -36.0}, {12.833, 13.5, -37.0}, {11.833, 12.833, -54.5}, {12.833, 14.167, -55.0}, {13.5, 14.167, -41.0}, {14.167, 14.917, -61.0}, {12.833, 14.167, -63.0}, {14.167, 14.917, -65.5}, {12.833, 14.167, -67.5}}},
	{Abbr: "Cep", LatinName: "Cepheus", EnglishName: "King of Ethiopia", ImageID: "cepheus", Boundaries: []Boundary{{0.0, 5.0, 80.0}, {0.0, 8.0, 85.0}, {20.167, 21.0, 80.0}, {20.167, 20.667, 75.0}, {20.6, 21.967, 54.833}, {21.967, 22.133, 52.75}, {20.537, 20.6, 60.917}, {20.0, 20.537, 59.5}, {0.0, 3.1, 66.0}, {23.583, 24.0, 63.0}, {0.0, 1.7, 54.0}}},
	{Abbr: "Cet", LatinName: "Cetus", EnglishName: "Whale", ImageID: "cetus", Boundaries: []Boundary{{2.167, 2.5, -7.5}, {2.5, 3.0, -10.0}, {0.333, 2.167, -10.5}, {1.667, 2.167, -12.0}, {0.0, 1.667, -14.5}, {23.833, 24.0, -14.5}, {0.0, 2.0, -24.0}, {23.833, 24.0, -24.0}}},
	{Abbr: "Cha", LatinName: "Chamaeleon", EnglishName: "Chameleon", ImageID: "chamaeleon", Boundaries: []Boundary{{12.833, 14.167, -79.0}, {14.167, 15.333, -80.0}, {11.833, 12.833, -81.5}, {12.833, 14.167, -83.0}, {11.833, 14.167, -85.0}, {11.833, 14.167, -88.0}}},
	{Abbr: "Cir", LatinName: "Circinus", EnglishName: "Compass", ImageID: "circinus", Boundaries: []Boundary{{14.167, 14.917, -65.5}, {14.917, 15.333, -67.0}, {12.833, 14.167, -67.5}, {14.167, 15.333, -70.0}, {14.917, 15.333, -71.5}, {14.167, 15.333, -76.0}}},
	{Abbr: "Col", LatinName: "Columba", EnglishName: "Dove", ImageID: "columba", Boundaries: []Boundary{{6.0, 6.5, -42.0}, {6.5, 6.833, -43.75}}},
	{Abbr: "Com", LatinName: "Coma Berenices", EnglishName: "Berenice's Hair", ImageID: "coma-berenices", Boundaries: []Boundary{{12.0, 12.333, -5.5}}},
	{Abbr: "CrA", LatinName: "Corona Australis", EnglishName: "Southern Crown", ImageID: "corona-australis", Boundaries: []Boundary{{18.0, 18.333, -34.5}, {18.333, 19.167, -35.0}, {17.833, 18.0, -38.5}, {18.0, 18.5, -40.0}, {18.5, 19.0, -40.333}}},
	{Abbr: "CrB", LatinName: "Corona Borealis", EnglishName: "Northern Crown", ImageID: "coronaBorealis", Boundaries: []Boundary{{15.25, 15.75, 42.0}, {15.75, 15.917, 34.5}, {16.167, 16.333, 26.0}}},
	{Abbr: "Crv", LatinName: "Corvus", EnglishName: "Crow", ImageID: "corvus", Boundaries: []Boundary{{11.583, 11.833, -16.5}, {11.833, 12.833, -18.5}}},
	{Abbr: "Crt", LatinName: "Crater", EnglishName: "Cup", ImageID: "crater", Boundaries: []Boundary{{10.583, 10.833, -13.0}, {10.833, 11.583, -16.5}}},
	{Abbr: "Cru", LatinName: "Crux", EnglishName: "Southern Cross", ImageID: "crux", Boundaries: []Boundary{{11.833, 12.833, -60.0}, {11.333, 11.833, -72.5}}},
	{Abbr: "Cyg", LatinName: "Cygnus", EnglishName: "Swan", ImageID: "cygnus", Boundaries: []Boundary{{19.167, 19.4, 39.167}, {19.4, 19.6, 38.0}, {19.6, 20.0, 35.333}, {20.0, 20.2, 35.0}, {20.2, 20.567, 34.0}, {20.567, 20.617, 31.417}, {20.617, 21.5, 31.417}, {21.5, 21.733, 31.75}, {21.733, 22.0, 30.0}, {19.0, 19.167, 27.5}}},
	{Abbr: "Del", LatinName: "Delphinus", EnglishName: "Dolphin", ImageID: "delphinus", Boundaries: []Boundary{{20.2, 20.983, 17.167}, {20.983, 21.467, 15.0}, {20.833, 20.983, 2.0}}},
	{Abbr: "Dor", LatinName: "Dorado", EnglishName: "Swordfish", ImageID: "dorado", Boundaries: []Boundary{{4.5, 5.0, -57.5}, {5.0, 5.5, -57.5}, {5.5, 6.0, -57.5}, {6.0, 6.5, -58.5}, {4.5, 5.0, -59.5}, {5.0, 5.5, -60.0}, {5.5, 6.0, -61.0}, {6.0, 6.5, -62.0}, {5.5, 6.0, -64.0}, {4.5, 5.0, -64.5}, {5.0, 5.5, -65.0}, {4.5, 5.0, -68.5}, {3.5, 4.5, -70.0}, {3.5, 4.5, -75.5}, {4.5, 5.0, -75.5}, {3.5, 4.5, -79.5}, {4.5, 6.5, -83.5}, {4.5, 6.5, -85.0}, {4.5, 6.5, -87.0}}},
	{Abbr: "Dra", LatinName: "Draco", EnglishName: "Dragon", ImageID: "draco", Boundaries: []Boundary{{9.167, 11.333, 73.5}, {11.333, 12.0, 66.5}, {12.0, 13.5, 64.0}, {13.5, 14.417, 63.0}, {20.417, 20.667, 67.0}, {20.0, 20.417, 61.5}, {19.767, 20.0, 59.5}, {19.417, 19.767, 58.0}, {14.417, 19.417, 55.5}, {15.25, 15.75, 53.0}, {15.75, 17.0, 51.5}}},
	{Abbr: "Equ", LatinName: "Equuleus", EnglishName: "Little Horse", ImageID: "equuleus", Boundaries: []Boundary{{21.467, 21.733, 13.0}, {21.733, 21.867, 12.833}, {20.983, 21.333, 2.0}}},
	{Abbr: "Eri", LatinName: "Eridanus", EnglishName: "River", ImageID: "eridanus", Boundaries: []Boundary{{5.083, 5.833, -1.75}, {5.0, 5.083, -3.0}, {4.833, 5.0, -5.0}, {4.7, 4.833, -6.5}, {4.5, 4.7, -9.0}, {4.333, 4.5, -10.5}, {4.0, 4.333, -13.0}, {3.5, 4.0, -15.5}, {3.0, 3.5, -17.5}, {2.667, 3.0, -21.5}, {2.0, 2.667, -24.0}, {2.0, 2.667, -25.5}, {2.667, 3.0, -26.5}, {3.0, 3.5, -29.0}, {3.5, 4.333, -30.0}, {4.333, 4.5, -31.0}, {4.5, 5.0, -33.0}, {5.0, 5.5, -35.0}, {5.5, 6.0, -37.0}, {5.5, 6.0, -42.0}, {5.0, 5.5, -40.5}, {4.5, 5.0, -39.583}, {2.167, 3.0, -52.5}, {1.333, 2.167, -58.5}}},
	{Abbr: "For", LatinName: "Fornax", EnglishName: "Furnace", ImageID: "fornax", Boundaries: []Boundary{{2.667, 3.0, -32.0}, {3.0, 3.5, -34.5}, {3.5, 4.0, -35.5}, {4.0, 4.333, -36.333}, {4.333, 4.5, -38.333}}},
	{Abbr: "Gem", LatinName: "Gemini", EnglishName: "Twins", ImageID: "gemini", Boundaries: []Boundary{{6.217, 7.5, 32.0}, {7.5, 7.75, 31.5}, {7.75, 8.0, 30.75}, {7.75, 8.0, 28.5}, {7.5, 7.75, 20.0}, {6.0, 6.217, 29.5}, {5.917, 6.0, 12.0}, {5.833, 5.917, 11.0}, {7.0, 7.25, 12.5}, {7.0, 7.25, 9.0}, {7.25, 7.5, 9.0}, {8.0, 8.083, 24.5}, {8.0, 8.083, 15.5}}},
	{Abbr: "Gru", LatinName: "Grus", EnglishName: "Crane", ImageID: "grus", Boundaries: []Boundary{{21.333, 21.867, -45.0}, {21.333, 22.0, -48.167}, {22.0, 23.333, -54.0}}},
	{Abbr: "Her", LatinName: "Hercules", EnglishName: "Hercules", ImageID: "hercules", Boundaries: []Boundary{{15.75, 16.333, 42.0}, {16.333, 17.25, 40.0}, {17.25, 17.967, 33.5}, {17.967, 18.175, 31.0}, {18.175, 18.333, 32.083}, {16.533, 16.867, 15.0}, {17.167, 17.967, 21.5}, {17.967, 18.25, 19.5}, {18.25, 18.867, 18.167}, {18.25, 18.425, 16.0}}},
	{Abbr: "Hor", LatinName: "Horologium", EnglishName: "Clock", ImageID: "horologium", Boundaries: []Boundary{{3.417, 4.5, -44.5}, {3.0, 3.5, -53.167}, {3.5, 4.5, -54.0}, {4.5, 5.0, -48.0}, {1.333, 3.0, -68.0}, {0.0, 3.5, -75.0}, {0.0, 3.5, -78.5}, {0.0, 3.5, -80.0}, {0.0, 4.5, -84.5}, {0.0, 4.5, -85.5}, {0.0, 4.5, -87.5}}},
	{Abbr: "Hya", LatinName: "Hydra", EnglishName: "Sea Serpent", ImageID: "hydra", Boundaries: []Boundary{{8.333, 8.583, -9.5}, {8.583, 9.083, -10.0}, {7.5, 8.333, -11.0}, {9.083, 9.75, -11.667}, {9.75, 10.25, -13.0}, {8.333, 8.667, -15.0}, {8.667, 9.333, -15.667}, {9.333, 9.917, -17.0}, {9.917, 10.25, -18.25}, {10.25, 10.583, -19.0}, {10.583, 10.833, -20.5}, {10.833, 11.333, -21.5}, {11.333, 11.833, -22.333}, {11.833, 12.833, -22.5}, {12.833, 14.25, -24.0}, {12.833, 14.25, -26.5}, {12.833, 13.5, -30.0}, {13.5, 14.25, -30.0}, {12.833, 13.5, -35.0}, {12.833, 13.5, -37.0}}},
	{Abbr: "Hyi", LatinName: "Hydrus", EnglishName: "Water Snake", ImageID: "hydrus", Boundaries: []Boundary{}},
	{Abbr: "Ind", LatinName: "Indus", EnglishName: "Indian", ImageID: "indus", Boundaries: []Boundary{{20.0, 20.5, -54.5}, {20.5, 21.333, -56.5}, {21.333, 23.333, -60.0}, {20.0, 21.333, -74.0}}},
	{Abbr: "Lac", LatinName: "Lacerta", EnglishName: "Lizard", ImageID: "lacerta", Boundaries: []Boundary{{21.867, 22.0, 38.333}, {22.0, 22.75, 32.083}, {22.0, 22.75, 29.5}}},
	{Abbr: "Leo", LatinName: "Leo", EnglishName: "Lion", ImageID: "leo", Boundaries: []Boundary{{10.867, 11.0, 33.0}, {10.75, 11.0, 19.5}, {11.0, 11.533, 17.0}, {9.833, 10.25, 11.5}, {10.25, 10.583, 3.333}, {10.583, 10.833, 0.0}, {10.833, 11.867, -0.5}, {11.867, 12.0, 1.5}}},
	{Abbr: "LMi", LatinName: "Leo Minor", EnglishName: "Little Lion", ImageID: "leo-minor", Boundaries: []Boundary{{9.0, 9.083, 44.0}, {9.0, 9.25, 41.5}, {9.25, 9.583, 39.75}, {9.583, 10.0, 38.5}, {10.0, 10.117, 36.5}, {10.117, 10.867, 34.0}, {10.867, 11.0, 33.0}}},
	{Abbr: "Lep", LatinName: "Lepus", EnglishName: "Hare", ImageID: "lepus", Boundaries: []Boundary{{6.0, 6.583, -39.0}}},
	{Abbr: "Lib", LatinName: "Libra", EnglishName: "Scales", ImageID: "libra", Boundaries: []Boundary{{14.667, 15.0, -12.0}, {15.0, 15.917, -12.0}, {15.917, 16.267, -13.0}, {16.267, 16.375, -15.333}, {15.0, 15.333, -18.25}, {15.333, 15.667, -19.25}, {15.667, 15.917, -20.667}, {14.25, 14.917, -21.0}, {14.917, 15.0, -22.0}, {14.25, 14.917, -24.5}, {14.917, 15.333, -24.5}, {15.333, 15.667, -26.0}, {14.917, 15.667, -29.5}, {14.917, 15.0, -25.5}, {14.25, 14.917, -28.5}}},
	{Abbr: "Lup", LatinName: "Lupus", EnglishName: "Wolf", ImageID: "lupus", Boundaries: []Boundary{{14.167, 14.917, -36.75}, {14.917, 15.333, -38.5}, {15.333, 15.667, -40.0}, {15.667, 16.0, -41.167}, {16.0, 16.375, -41.5}, {16.375, 16.625, -42.333}, {14.167, 14.917, -43.0}, {14.917, 15.667, -45.5}, {14.167, 14.917, -51.5}, {14.917, 15.333, -54.0}, {14.167, 14.917, -55.0}}},
	{Abbr: "Lyn", LatinName: "Lynx", EnglishName: "Lynx", ImageID: "lynx", Boundaries: []Boundary{{6.1, 6.5, 54.0}, {6.5, 6.8, 50.0}, {6.8, 7.367, 47.5}, {7.367, 7.9, 44.5}, {7.9, 8.0, 39.5}, {8.0, 9.0, 35.5}, {9.0, 9.25, 36.5}, {9.25, 10.0, 30.0}, {9.0, 9.25, 31.5}}},
	{Abbr: "Lyr", LatinName: "Lyra", EnglishName: "Lyre", ImageID: "lyra", Boundaries: []Boundary{{18.175, 18.9, 36.0}, {18.9, 19.0, 35.5}, {18.333, 18.5, 32.0}, {18.5, 19.0, 30.0}}},
	{Abbr: "Men", LatinName: "Mensa", EnglishName: "Table Mountain", ImageID: "mensa", Boundaries: []Boundary{{0.0, 4.5, -85.5}, {4.5, 6.5, -87.0}, {0.0, 4.5, -87.5}}},
	{Abbr: "Mic", LatinName: "Microscopium", EnglishName: "Microscope", ImageID: "microscopium", Boundaries: []Boundary{{20.0, 20.333, -40.5}, {20.333, 21.333, -42.5}, {20.333, 21.333, -47.0}}},
	{Abbr: "Mon", LatinName: "Monoceros", EnglishName: "Unicorn", ImageID: "monoceros", Boundaries: []Boundary{{6.667, 7.0, 5.5}, {6.367, 6.667, 1.0}, {7.25, 7.5, 0.5}, {6.667, 7.25, -1.0}, {6.0, 6.367, -3.0}, {6.0, 6.5, -5.0}, {6.5, 6.833, -7.0}, {6.833, 7.5, -8.333}, {7.5, 8.333, -6.0}}},
	{Abbr: "Mus", LatinName: "Musca", EnglishName: "Fly", ImageID: "musca", Boundaries: []Boundary{{11.333, 11.833, -70.0}, {10.583, 11.333, -72.0}, {11.333, 11.833, -72.5}, {11.833, 12.833, -74.0}, {12.833, 14.167, -74.5}, {11.833, 12.833, -77.5}}},
	{Abbr: "Nor", LatinName: "Norma", EnglishName: "Carpenter's Square", ImageID: "norma", Boundaries: []Boundary{{15.667, 16.0, -52.0}, {15.333, 15.667, -51.0}, {16.0, 16.625, -48.75}, {16.0, 16.375, -47.5}, {15.667, 16.0, -47.0}, {14.917, 15.333, -54.0}, {15.333, 16.0, -56.0}, {15.333, 16.0, -61.333}, {16.0, 16.5, -66.5}}},
	{Abbr: "Oct", LatinName: "Octans", EnglishName: "Octant", ImageID: "octans", Boundaries: []Boundary{{17.5, 21.333, -83.5}, {21.333, 23.333, -83.5}, {23.333, 24.0, -86.0}, {17.5, 23.333, -88.0}, {0.0, 9.333, -90.0}, {17.5, 24.0, -90.0}}},
	{Abbr: "Oph", LatinName: "Ophiuchus", EnglishName: "Serpent Bearer", ImageID: "ophiuchus", Boundaries: []Boundary{{16.375, 16.625, -24.583}, {16.625, 17.0, -25.5}, {17.0, 17.5, -25.5}, {17.5, 17.833, -26.0}, {17.833, 18.0, -28.5}, {16.267, 17.0, -8.0}, {17.0, 17.167, -4.0}, {16.867, 17.583, 7.0}, {17.583, 17.667, 7.5}, {17.833, 17.967, -2.0}, {17.167, 17.833, -3.25}, {17.833, 18.25, 9.5}, {18.25, 18.867, 12.5}, {16.375, 16.625, -16.5}}},
	{Abbr: "Ori", LatinName: "Orion", EnglishName: "Hunter", ImageID: "orion", Boundaries: []Boundary{{6.0, 6.367, 17.5}, {5.917, 6.0, 12.0}, {5.5, 5.833, 3.0}, {5.833, 6.0, 2.0}, {5.0, 5.083, 6.0}, {6.0, 6.367, 6.0}, {5.5, 5.833, -4.0}, {5.833, 6.0, -1.75}, {5.083, 5.833, -1.75}, {5.833, 6.0, 0.0}, {5.0, 5.083, 12.5}}},
	{Abbr: "Pav", LatinName: "Pavo", EnglishName: "Peacock", ImageID: "pavo", Boundaries: []Boundary{{19.0, 19.333, -58.5}, {19.333, 19.667, -59.167}, {18.0, 18.5, -60.0}, {19.667, 20.0, -61.5}, {18.5, 19.0, -63.5}, {19.0, 19.5, -64.5}, {20.0, 20.5, -64.0}, {19.5, 20.0, -65.5}, {18.0, 18.5, -65.5}, {17.5, 18.0, -75.0}, {18.0, 20.0, -70.0}, {18.0, 21.333, -80.0}}},
	{Abbr: "Peg", LatinName: "Pegasus", EnglishName: "Winged Horse", ImageID: "pegasus", Boundaries: []Boundary{{22.0, 22.75, 29.5}, {22.75, 23.75, 29.5}, {21.867, 22.0, 12.833}, {22.0, 22.75, 12.167}, {21.333, 21.867, 1.0}, {21.867, 22.75, 0.0}}},
	{Abbr: "Per", LatinName: "Perseus", EnglishName: "Hero", ImageID: "perseus", Boundaries: []Boundary{{0.867, 1.1, 50.5}, {1.1, 1.8, 45.0}, {0.867, 1.1, 43.0}, {0.0, 2.5, 40.5}, {2.5, 2.567, 40.5}, {1.8, 2.0, 38.5}, {2.567, 2.717, 38.5}, {2.717, 3.5, 36.75}}},
	{Abbr: "Phe", LatinName: "Phoenix", EnglishName: "Phoenix", ImageID: "phoenix", Boundaries: []Boundary{{0.0, 1.333, -46.5}, {23.333, 24.0, -46.5}, {1.333, 2.167, -49.75}, {0.0, 1.333, -53.5}, {23.333, 24.0, -53.5}, {0.0, 1.333, -58.5}, {23.333, 24.0, -58.5}}},
	{Abbr: "Pic", LatinName: "Pictor", EnglishName: "Painter's Easel", ImageID: "pictor", Boundaries: []Boundary{{6.5, 6.833, -48.0}, {6.0, 6.5, -49.75}, {5.0, 5.5, -50.75}, {6.5, 6.833, -51.0}, {5.5, 6.0, -52.5}, {6.0, 6.167, -53.5}, {6.167, 6.5, -54.5}, {4.5, 5.0, -55.0}, {5.0, 5.5, -55.0}, {6.833, 7.167, -62.0}, {7.167, 7.5, -63.5}, {4.5, 5.0, -64.5}, {5.0, 5.5, -71.5}, {5.5, 6.0, -72.5}, {4.5, 5.0, -71.0}, {5.0, 6.0, -76.0}, {4.5, 5.5, -80.0}}},
	{Abbr: "Psc", LatinName: "Pisces", EnglishName: "Fish", ImageID: "pisces", Boundaries: []Boundary{{22.75, 23.833, 3.5}, {23.833, 24.0, 2.0}, {0.0, 2.0, 0.0}, {0.0, 0.333, -10.0}, {23.833, 24.0, -10.0}}},
	{Abbr: "PsA", LatinName: "Piscis Austrinus", EnglishName: "Southern Fish", ImageID: "piscis-austrinus", Boundaries: []Boundary{{23.0, 23.333, -25.5}, {21.333, 21.867, -36.0}, {21.867, 23.0, -37.5}}},
	{Abbr: "Pup", LatinName: "Puppis", EnglishName: "Stern", ImageID: "puppis", Boundaries: []Boundary{{7.667, 8.333, -41.25}, {7.667, 8.333, -44.333}, {6.833, 7.667, -46.0}, {6.833, 7.5, -51.5}, {8.333, 8.583, -49.0}, {8.583, 9.333, -49.5}, {8.333, 8.583, -52.5}}},
	{Abbr: "Pyx", LatinName: "Pyxis", EnglishName: "Compass", ImageID: "pyxis", Boundaries: []Boundary{{8.333, 9.333, -38.333}}},
	{Abbr: "Ret", LatinName: "Reticulum", EnglishName: "Net", ImageID: "reticulum", Boundaries: []Boundary{{3.5, 4.5, -66.5}, {3.0, 3.5, -69.75}}},
	{Abbr: "Sge", LatinName: "Sagitta", EnglishName: "Arrow", ImageID: "sagitta", Boundaries: []Boundary{{19.25, 19.4, 20.5}, {19.4, 19.667, 20.0}, {19.667, 20.0, 19.167}, {20.0, 20.2, 17.5}, {19.25, 19.833, 16.167}}},
	{Abbr: "Sgr", LatinName: "Sagittarius", EnglishName: "Archer", ImageID: "sagittarius", Boundaries: []Boundary{{19.167, 19.5, -29.0}, {18.0, 18.333, -30.0}, {19.5, 19.833, -30.0}, {18.333, 18.5, -30.0}, {19.833, 20.0, -31.0}, {18.5, 19.167, -31.333}, {20.0, 20.333, -40.5}, {19.5, 20.0, -42.0}, {19.5, 20.333, -45.0}}},
	{Abbr: "Sco", LatinName: "Scorpius", EnglishName: "Scorpion", ImageID: "scorpius", Boundaries: []Boundary{{16.267, 16.375, -19.25}, {15.667, 16.0, -27.833}, {16.0, 16.375, -28.0}, {16.375, 16.5, -29.833}, {16.5, 17.0, -30.0}, {17.0, 17.5, -31.167}, {17.5, 17.833, -37.5}, {17.833, 18.0, -38.5}, {17.0, 17.5, -37.0}, {17.5, 18.0, -46.5}, {17.0, 17.5, -44.5}}},
	{Abbr: "Scl", LatinName: "Sculptor", EnglishName: "Sculptor", ImageID: "sculptor", Boundaries: []Boundary{{23.333, 23.833, -27.0}, {0.0, 2.667, -30.0}, {23.833, 24.0, -30.0}, {21.867, 23.0, -37.5}, {0.0, 1.333, -39.75}, {23.333, 24.0, -39.75}}},
	{Abbr: "Sct", LatinName: "Scutum", EnglishName: "Shield", ImageID: "scutum", Boundaries: []Boundary{{18.175, 18.25, -11.333}, {18.25, 18.667, -11.333}, {18.667, 18.867, -12.0}, {18.867, 19.25, -16.0}}},
	{Abbr: "Ser", LatinName: "Serpens", EnglishName: "Serpent", ImageID: "serpens-cauda", Boundaries: []Boundary{{15.917, 16.167, 26.0}, {15.333, 15.75, 22.5}, {15.75, 16.0, 16.0}, {16.0, 16.533, 15.0}, {15.0, 15.917, -12.0}, {17.0, 17.167, -4.0}, {17.167, 17.583, 4.0}, {17.833, 17.967, 4.5}, {17.967, 18.25, 1.0}, {18.25, 18.867, 2.0}}},
	{Abbr: "Sex", LatinName: "Sextans", EnglishName: "Sextant", ImageID: "sextans", Boundaries: []Boundary{{10.833, 11.333, -10.0}, {10.25, 10.583, -11.0}}},
	{Abbr: "Tau", LatinName: "Taurus", EnglishName: "Bull", ImageID: "taurus", Boundaries: []Boundary{{4.5, 4.667, 24.5}, {5.0, 5.167, 24.0}, {4.667, 5.0, 23.75}, {5.167, 5.317, 19.0}, {5.317, 5.5, 14.0}, {5.5, 5.583, 12.0}, {5.583, 5.917, 11.0}, {5.917, 6.0, 9.0}, {5.7, 5.883, 26.5}}},
	{Abbr: "Tel", LatinName: "Telescopium", EnglishName: "Telescope", ImageID: "telescopium", Boundaries: []Boundary{{18.5, 19.0, -49.0}, {19.0, 19.5, -50.0}, {19.5, 20.0, -51.5}}},
	{Abbr: "Tri", LatinName: "Triangulum", EnglishName: "Triangle", ImageID: "triangulum", Boundaries: []Boundary{{1.917, 2.0, 27.25}, {2.0, 2.5, 26.5}, {2.5, 2.717, 22.833}}},
	{Abbr: "TrA", LatinName: "Triangulum Australe", EnglishName: "Southern Triangle", ImageID: "triangulum-australe", Boundaries: []Boundary{{15.333, 16.0, -61.333}, {16.0, 16.625, -62.0}, {16.625, 17.0, -63.5}, {17.0, 17.5, -64.0}, {15.333, 16.0, -65.5}, {16.0, 16.5, -66.5}, {15.333, 16.5, -72.5}}},
	{Abbr: "Tuc", LatinName: "Tucana", EnglishName: "Toucan", ImageID: "tucana", Boundaries: []Boundary{{21.333, 23.333, -60.0}, {0.0, 1.333, -64.5}, {23.333, 24.0, -64.5}, {0.0, 1.333, -70.0}, {23.333, 24.0, -70.0}, {21.333, 23.333, -70.0}, {21.333, 23.333, -78.0}}},
	{Abbr: "UMa", LatinName: "Ursa Major", EnglishName: "Great Bear", ImageID: "ursa-major", Boundaries: []Boundary{{7.967, 8.417, 60.0}, {8.417, 9.083, 50.0}, {9.083, 10.167, 47.0}, {10.167, 10.783, 47.0}, {10.783, 11.0, 44.0}, {11.0, 12.0, 40.0}, {12.083, 13.5, 53.0}, {14.033, 14.417, 55.5}}},
	{Abbr: "UMi", LatinName: "Ursa Minor", EnglishName: "Little Bear", ImageID: "ursa-minor", Boundaries: []Boundary{{0.0, 24.0, 88.0}, {8.0, 14.5, 86.5}, {21.0, 23.0, 86.167}, {18.0, 21.0, 86.0}, {10.667, 14.5, 80.0}, {17.5, 18.0, 80.0}, {16.533, 17.5, 75.0}, {13.0, 16.533, 70.0}, {14.033, 15.667, 66.0}}},
	{Abbr: "Vel", LatinName: "Vela", EnglishName: "Sails", ImageID: "vela", Boundaries: []Boundary{{9.333, 10.75, -46.5}, {8.333, 8.583, -55.5}, {8.583, 9.333, -49.5}, {8.333, 8.583, -49.0}}},
	{Abbr: "Vir", LatinName: "Virgo", EnglishName: "Virgin", ImageID: "virgo", Boundaries: []Boundary{{10.833, 11.867, -0.5}, {12.0, 12.333, -5.5}, {12.333, 12.833, -7.0}, {12.833, 14.25, -8.5}, {14.25, 14.667, -8.0}, {12.833, 14.25, -19.5}}},
	{Abbr: "Vol", LatinName: "Volans", EnglishName: "Flying Fish", ImageID: "volans", Boundaries: []Boundary{{6.5, 7.667, -73.5}, {7.667, 8.333, -76.5}, {6.5, 7.667, -77.0}, {8.333, 9.333, -77.5}, {6.0, 6.5, -78.5}, {6.5, 9.333, -81.5}, {6.5, 9.333, -83.5}, {6.5, 9.333, -85.5}}},
	{Abbr: "Vul", LatinName: "Vulpecula", EnglishName: "Fox", ImageID: "vulpecula", Boundaries: []Boundary{{19.167, 19.4, 27.5}, {19.4, 19.667, 27.0}, {19.667, 20.0, 24.667}, {20.0, 20.2, 24.0}, {20.2, 20.567, 23.5}, {20.567, 20.983, 22.0}}},
}
