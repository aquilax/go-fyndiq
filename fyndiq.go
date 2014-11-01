// Package fyndiq prvides library for inteaction with Fyndiq's API
package fyndiq

const (
	apiURL     = "https://fyndiq.se/"
	apiSegment = "api"
	apiVersion = "v1"
)

// FyndiqAPI represents the main Api type
type FyndiqAPI struct {
	user     string
	token    string
	testMode bool
}

// NewFyndiqAPI creates new API Instance
func NewFyndiqAPI(user string, token string, testMode bool) *FyndiqAPI {
	return &FyndiqAPI{user, token, testMode}
}
