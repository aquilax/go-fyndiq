package fyndiq

import (
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
