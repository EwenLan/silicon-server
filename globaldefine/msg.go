package globaldefine

// VersionInfoPrototype
type VersionInfoPrototype struct {
	BaseGoVersion   string
	SoftwareVersion string
	ProjectHome     string
	Author          string
	Email           string
	BuildDate       string
}

type ProcessorReusingTestRequest struct {
	MessageA string
	MessageB string
	MessageC string
}
type ProcessorReusingTestResponse struct {
	ResponseA string
	ResponseB string
	ResponseC string
}

type CalculatingRequest struct {
	NumA string
	NumB string
	Op   string
}

type CalculatingResponse struct {
	Ans string
}
