package fyndiq

const (
	apiURL     = "https://fyndiq.se/"
	apiSegment = "api"
	apiVersion = "v1"
)

// FyndiqAPI represents the main Api type
type FyndiqAPI struct {
	user  string
	token string
}

// NewFyndiqAPI creates new API Instance
func NewFyndiqAPI(user string, token string) *FyndiqAPI {
	return &FyndiqAPI{user, token}
}
