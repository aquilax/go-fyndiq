package fyndiq

// Generic meta data response header
type MetaData struct {
	Limit      int    `json: "limit"`
	Next       string `json: "next"`
	Offset     int    `json: "offset"`
	Previous   string `json: "previous"`
	TotalCount int    `json: "total_count"`
}

// Generic response with meta data
type MetaResponse struct {
	Meta MetaData `json: "meta"`
}
