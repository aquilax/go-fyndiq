package fyndiq

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	separatorPath = "/"
)

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

func httpRequest(method string, url string, body io.Reader) (*response, error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	return &response{
		resp.StatusCode,
		contents,
		resp.Header,
	}, err
}
