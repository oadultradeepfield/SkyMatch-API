package nova

type LoginResponse struct {
	Status  string `json:"status"`
	Session string `json:"session"`
	Message string `json:"message"`
}

type UploadResponse struct {
	Status string `json:"status"`
	SubID  int    `json:"subid"`
}

type Submission struct {
	Jobs []int `json:"jobs"`
}

type JobStatusResponse struct {
	Status string `json:"status"`
}

type JobInfo struct {
	Calibration    Calibration `json:"calibration"`
	ObjectsInField []string    `json:"objects_in_field"`
}

type Calibration struct {
	RA  float64 `json:"ra"`
	Dec float64 `json:"dec"`
}

type Annotation struct {
	Names  []string `json:"names"`
	PixelX float64  `json:"pixelx"`
	PixelY float64  `json:"pixely"`
	Type   string   `json:"type"`
}
