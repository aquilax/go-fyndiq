package fyndiq

// RequestParams is used for passing query params
type RequestParams map[string]string

type response struct {
	code int
	body []byte
	header map[string][]string
}

// MetaData holds the generic meta data response header
type MetaData struct {
	Limit      int    `json: "limit"`
	Next       string `json: "next"`
	Offset     int    `json: "offset"`
	Previous   string `json: "previous"`
	TotalCount int    `json: "total_count"`
}

// MetaResponse holds generic response with meta data
type MetaResponse struct {
	Meta MetaData `json: "meta"`
}
