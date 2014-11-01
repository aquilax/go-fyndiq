package fyndiq

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	separatorPath = "/"
)

type RequestParams map[string]string

func getPath(segments []string) string {
	segments = append([]string{apiSegment, apiVersion}, segments...)
	return strings.Join(segments, separatorPath)
}

func (fapi *FyndiqAPI) getURL(path string, params RequestParams) string {
	parsedURL, _ := url.Parse(apiURL)
	values := url.Values{}
	values.Set("user", fapi.user)
	values.Set("token", fapi.token)
	for name, value := range params {
		values.Set(name, value)
	}
	parsedURL.Path = path
	parsedURL.RawQuery = values.Encode()
	return parsedURL.String()
}

func getJSONBlob(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
