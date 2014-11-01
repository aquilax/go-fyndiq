package fyndiq

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	SEPARATOR_PATH = "/"
)

type RequestParams map[string]string

func getPath(segments []string) string {
	segments = append([]string{API_SEGMENT, API_VERSION}, segments...)
	return strings.Join(segments, SEPARATOR_PATH)
}

func (fapi *FyndiqApi) getUrl(path string, params RequestParams) string {
	parsedUrl, _ := url.Parse(URL)
	values := url.Values{}
	values.Set("user", fapi.user)
	values.Set("token", fapi.token)
	for name, value := range params {
		values.Set(name, value)
	}
	parsedUrl.Path = path
	parsedUrl.RawQuery = values.Encode()
	return parsedUrl.String()
}

func getJsonBlob(url string) ([]byte, error) {
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
