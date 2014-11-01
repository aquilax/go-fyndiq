package fyndiq

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type RequestParams map[string]string

func (fapi *FyndiqApi) getUrl(path string, params RequestParams) string {
	parsedUrl, _ := url.Parse(URL)
	values := url.Values{}
	values.Set("user", fapi.user)
	values.Set("token", fapi.token)
	for name, value := range params {
		values.Set(name, value)
	}
	parsedUrl.Path = parsedUrl.Path + path
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
