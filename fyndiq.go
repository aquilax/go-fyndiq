package fyndiq

const (
	apiURL     = "https://fyndiq.se/"
	apiSegment = "api"
	apiVersion = "v1"
)

// Main Api type
type FyndiqAPI struct {
	user  string
	token string
}

// Creates new API Instance
func NewFyndiqAPI(user string, token string) *FyndiqAPI {
	return &FyndiqAPI{user, token}
}
