package model

import (
	"bufio"
	"bytes"
	_ "embed"
	"strconv"
	"strings"
)

const (
	noirLabBaseURL = "https://storage.noirlab.edu/media/archives/images/thumb350x/"
	noirLabSuffix  = "-ann.jpg"
)

//go:embed data/constellation_boundaries.dat
var boundaryData []byte

type Constellation struct {
	Abbr        string
	LatinName   string
	EnglishName string
	ImageID     string
}

type boundaryRecord struct {
	RALow, RAHigh, DecLow float64
	Abbr                  string
}

var boundaries []boundaryRecord

func init() {
	boundaries = parseBoundaryData(boundaryData)
}

func parseBoundaryData(data []byte) []boundaryRecord {
	var records []boundaryRecord
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		raLow, err1 := strconv.ParseFloat(fields[0], 64)
		raHigh, err2 := strconv.ParseFloat(fields[1], 64)
		decLow, err3 := strconv.ParseFloat(fields[2], 64)
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}
		abbr := fields[3]
		records = append(records, boundaryRecord{
			RALow:  raLow,
			RAHigh: raHigh,
			DecLow: decLow,
			Abbr:   abbr,
		})
	}
	return records
}

func (c *Constellation) ImageURL() string {
	return noirLabBaseURL + c.ImageID + noirLabSuffix
}

func getConstellationByAbbr(abbr string) *Constellation {
	for i := range Constellations {
		if Constellations[i].Abbr == abbr {
			return &Constellations[i]
		}
	}
	return nil
}

func GetConstellationByCoords(ra, dec float64) *Constellation {
	for ra < 0 {
		ra += 360
	}
	for ra >= 360 {
		ra -= 360
	}

	ra1875, dec1875 := precessJ2000ToB1875(ra, dec)
	raHours := ra1875 / 15.0

	for _, b := range boundaries {
		if dec1875 >= b.DecLow && raHours >= b.RALow && raHours < b.RAHigh {
			return getConstellationByAbbr(b.Abbr)
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
	{Abbr: "And", LatinName: "Andromeda", EnglishName: "Princess of Ethiopia", ImageID: "andromeda"},
	{Abbr: "Ant", LatinName: "Antlia", EnglishName: "Air Pump", ImageID: "antlia"},
	{Abbr: "Aps", LatinName: "Apus", EnglishName: "Bird of Paradise", ImageID: "apus"},
	{Abbr: "Aqr", LatinName: "Aquarius", EnglishName: "Water Bearer", ImageID: "aquarius"},
	{Abbr: "Aql", LatinName: "Aquila", EnglishName: "Eagle", ImageID: "aquila"},
	{Abbr: "Ara", LatinName: "Ara", EnglishName: "Altar", ImageID: "ara"},
	{Abbr: "Ari", LatinName: "Aries", EnglishName: "Ram", ImageID: "aries"},
	{Abbr: "Aur", LatinName: "Auriga", EnglishName: "Charioteer", ImageID: "auriga"},
	{Abbr: "Boo", LatinName: "Bootes", EnglishName: "Herdsman", ImageID: "bootes"},
	{Abbr: "Cae", LatinName: "Caelum", EnglishName: "Chisel", ImageID: "caelum"},
	{Abbr: "Cam", LatinName: "Camelopardalis", EnglishName: "Giraffe", ImageID: "camelopardalis"},
	{Abbr: "Cnc", LatinName: "Cancer", EnglishName: "Crab", ImageID: "cancer"},
	{Abbr: "CVn", LatinName: "Canes Venatici", EnglishName: "Hunting Dogs", ImageID: "canesVenatici"},
	{Abbr: "CMa", LatinName: "Canis Major", EnglishName: "Great Dog", ImageID: "canis-major"},
	{Abbr: "CMi", LatinName: "Canis Minor", EnglishName: "Little Dog", ImageID: "canis-minor"},
	{Abbr: "Cap", LatinName: "Capricornus", EnglishName: "Sea Goat", ImageID: "capricornus"},
	{Abbr: "Car", LatinName: "Carina", EnglishName: "Keel", ImageID: "carina"},
	{Abbr: "Cas", LatinName: "Cassiopeia", EnglishName: "Queen of Ethiopia", ImageID: "cassiopeia"},
	{Abbr: "Cen", LatinName: "Centaurus", EnglishName: "Centaur", ImageID: "centaurus"},
	{Abbr: "Cep", LatinName: "Cepheus", EnglishName: "King of Ethiopia", ImageID: "cepheus"},
	{Abbr: "Cet", LatinName: "Cetus", EnglishName: "Whale", ImageID: "cetus"},
	{Abbr: "Cha", LatinName: "Chamaeleon", EnglishName: "Chameleon", ImageID: "chamaeleon"},
	{Abbr: "Cir", LatinName: "Circinus", EnglishName: "Compass", ImageID: "circinus"},
	{Abbr: "Col", LatinName: "Columba", EnglishName: "Dove", ImageID: "columba"},
	{Abbr: "Com", LatinName: "Coma Berenices", EnglishName: "Berenice's Hair", ImageID: "coma-berenices"},
	{Abbr: "CrA", LatinName: "Corona Australis", EnglishName: "Southern Crown", ImageID: "corona-australis"},
	{Abbr: "CrB", LatinName: "Corona Borealis", EnglishName: "Northern Crown", ImageID: "coronaBorealis"},
	{Abbr: "Crv", LatinName: "Corvus", EnglishName: "Crow", ImageID: "corvus"},
	{Abbr: "Crt", LatinName: "Crater", EnglishName: "Cup", ImageID: "crater"},
	{Abbr: "Cru", LatinName: "Crux", EnglishName: "Southern Cross", ImageID: "crux"},
	{Abbr: "Cyg", LatinName: "Cygnus", EnglishName: "Swan", ImageID: "cygnus"},
	{Abbr: "Del", LatinName: "Delphinus", EnglishName: "Dolphin", ImageID: "delphinus"},
	{Abbr: "Dor", LatinName: "Dorado", EnglishName: "Swordfish", ImageID: "dorado"},
	{Abbr: "Dra", LatinName: "Draco", EnglishName: "Dragon", ImageID: "draco"},
	{Abbr: "Equ", LatinName: "Equuleus", EnglishName: "Little Horse", ImageID: "equuleus"},
	{Abbr: "Eri", LatinName: "Eridanus", EnglishName: "River", ImageID: "eridanus"},
	{Abbr: "For", LatinName: "Fornax", EnglishName: "Furnace", ImageID: "fornax"},
	{Abbr: "Gem", LatinName: "Gemini", EnglishName: "Twins", ImageID: "gemini"},
	{Abbr: "Gru", LatinName: "Grus", EnglishName: "Crane", ImageID: "grus"},
	{Abbr: "Her", LatinName: "Hercules", EnglishName: "Hercules", ImageID: "hercules"},
	{Abbr: "Hor", LatinName: "Horologium", EnglishName: "Clock", ImageID: "horologium"},
	{Abbr: "Hya", LatinName: "Hydra", EnglishName: "Sea Serpent", ImageID: "hydra"},
	{Abbr: "Hyi", LatinName: "Hydrus", EnglishName: "Water Snake", ImageID: "hydrus"},
	{Abbr: "Ind", LatinName: "Indus", EnglishName: "Indian", ImageID: "indus"},
	{Abbr: "Lac", LatinName: "Lacerta", EnglishName: "Lizard", ImageID: "lacerta"},
	{Abbr: "Leo", LatinName: "Leo", EnglishName: "Lion", ImageID: "leo"},
	{Abbr: "LMi", LatinName: "Leo Minor", EnglishName: "Little Lion", ImageID: "leo-minor"},
	{Abbr: "Lep", LatinName: "Lepus", EnglishName: "Hare", ImageID: "lepus"},
	{Abbr: "Lib", LatinName: "Libra", EnglishName: "Scales", ImageID: "libra"},
	{Abbr: "Lup", LatinName: "Lupus", EnglishName: "Wolf", ImageID: "lupus"},
	{Abbr: "Lyn", LatinName: "Lynx", EnglishName: "Lynx", ImageID: "lynx"},
	{Abbr: "Lyr", LatinName: "Lyra", EnglishName: "Lyre", ImageID: "lyra"},
	{Abbr: "Men", LatinName: "Mensa", EnglishName: "Table Mountain", ImageID: "mensa"},
	{Abbr: "Mic", LatinName: "Microscopium", EnglishName: "Microscope", ImageID: "microscopium"},
	{Abbr: "Mon", LatinName: "Monoceros", EnglishName: "Unicorn", ImageID: "monoceros"},
	{Abbr: "Mus", LatinName: "Musca", EnglishName: "Fly", ImageID: "musca"},
	{Abbr: "Nor", LatinName: "Norma", EnglishName: "Carpenter's Square", ImageID: "norma"},
	{Abbr: "Oct", LatinName: "Octans", EnglishName: "Octant", ImageID: "octans"},
	{Abbr: "Oph", LatinName: "Ophiuchus", EnglishName: "Serpent Bearer", ImageID: "ophiuchus"},
	{Abbr: "Ori", LatinName: "Orion", EnglishName: "Hunter", ImageID: "orion"},
	{Abbr: "Pav", LatinName: "Pavo", EnglishName: "Peacock", ImageID: "pavo"},
	{Abbr: "Peg", LatinName: "Pegasus", EnglishName: "Winged Horse", ImageID: "pegasus"},
	{Abbr: "Per", LatinName: "Perseus", EnglishName: "Hero", ImageID: "perseus"},
	{Abbr: "Phe", LatinName: "Phoenix", EnglishName: "Phoenix", ImageID: "phoenix"},
	{Abbr: "Pic", LatinName: "Pictor", EnglishName: "Painter's Easel", ImageID: "pictor"},
	{Abbr: "Psc", LatinName: "Pisces", EnglishName: "Fish", ImageID: "pisces"},
	{Abbr: "PsA", LatinName: "Piscis Austrinus", EnglishName: "Southern Fish", ImageID: "piscis-austrinus"},
	{Abbr: "Pup", LatinName: "Puppis", EnglishName: "Stern", ImageID: "puppis"},
	{Abbr: "Pyx", LatinName: "Pyxis", EnglishName: "Compass", ImageID: "pyxis"},
	{Abbr: "Ret", LatinName: "Reticulum", EnglishName: "Net", ImageID: "reticulum"},
	{Abbr: "Sge", LatinName: "Sagitta", EnglishName: "Arrow", ImageID: "sagitta"},
	{Abbr: "Sgr", LatinName: "Sagittarius", EnglishName: "Archer", ImageID: "sagittarius"},
	{Abbr: "Sco", LatinName: "Scorpius", EnglishName: "Scorpion", ImageID: "scorpius"},
	{Abbr: "Scl", LatinName: "Sculptor", EnglishName: "Sculptor", ImageID: "sculptor"},
	{Abbr: "Sct", LatinName: "Scutum", EnglishName: "Shield", ImageID: "scutum"},
	{Abbr: "Ser", LatinName: "Serpens", EnglishName: "Serpent", ImageID: "serpens-cauda"},
	{Abbr: "Sex", LatinName: "Sextans", EnglishName: "Sextant", ImageID: "sextans"},
	{Abbr: "Tau", LatinName: "Taurus", EnglishName: "Bull", ImageID: "taurus"},
	{Abbr: "Tel", LatinName: "Telescopium", EnglishName: "Telescope", ImageID: "telescopium"},
	{Abbr: "Tri", LatinName: "Triangulum", EnglishName: "Triangle", ImageID: "triangulum"},
	{Abbr: "TrA", LatinName: "Triangulum Australe", EnglishName: "Southern Triangle", ImageID: "triangulum-australe"},
	{Abbr: "Tuc", LatinName: "Tucana", EnglishName: "Toucan", ImageID: "tucana"},
	{Abbr: "UMa", LatinName: "Ursa Major", EnglishName: "Great Bear", ImageID: "ursa-major"},
	{Abbr: "UMi", LatinName: "Ursa Minor", EnglishName: "Little Bear", ImageID: "ursa-minor"},
	{Abbr: "Vel", LatinName: "Vela", EnglishName: "Sails", ImageID: "vela"},
	{Abbr: "Vir", LatinName: "Virgo", EnglishName: "Virgin", ImageID: "virgo"},
	{Abbr: "Vol", LatinName: "Volans", EnglishName: "Flying Fish", ImageID: "volans"},
	{Abbr: "Vul", LatinName: "Vulpecula", EnglishName: "Fox", ImageID: "vulpecula"},
}
