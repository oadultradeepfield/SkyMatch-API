package model

type JobStatus string

const (
	StatusQueued             JobStatus = "QUEUED"
	StatusIdentifyingObjects JobStatus = "IDENTIFYING_OBJECTS"
	StatusGettingMoreDetails JobStatus = "GETTING_MORE_DETAILS"
	StatusSuccess            JobStatus = "SUCCESS"
	StatusFailure            JobStatus = "FAILURE"
	StatusCancelled          JobStatus = "CANCELLED"
)

type ObjectType string

const (
	ObjectTypeStar ObjectType = "STAR"
	ObjectTypeDSO  ObjectType = "DEEP_SKY_OBJECT"
)

type DeepSkyObjectType string

const (
	DSOOpenCluster     DeepSkyObjectType = "OPEN_CLUSTER"
	DSOGlobularCluster DeepSkyObjectType = "GLOBULAR_CLUSTER"
	DSOGalaxy          DeepSkyObjectType = "GALAXY"
	DSONebula          DeepSkyObjectType = "NEBULA"
	DSOSupernova       DeepSkyObjectType = "SUPERNOVA"
)

type SpectralClass string

const (
	SpectralO SpectralClass = "O"
	SpectralB SpectralClass = "B"
	SpectralA SpectralClass = "A"
	SpectralF SpectralClass = "F"
	SpectralG SpectralClass = "G"
	SpectralK SpectralClass = "K"
	SpectralM SpectralClass = "M"
)

type IdentifiedObject struct {
	Type            ObjectType
	Identifier      string
	Name            string
	Constellation   *Constellation
	XCoordinate     float64
	YCoordinate     float64
	VMagnitude      *float64
	SpectralClass   SpectralClass
	DistanceParsecs *float64
	DSOType         DeepSkyObjectType
}

type SolveResult struct {
	JobID             string
	Status            JobStatus
	AnnotatedImageURL string
	Objects           []IdentifiedObject
	NovaJobID         int
}
