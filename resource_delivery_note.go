package fyndiq

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	deliveryNotesSegment = "hamta-plocklistor"
)

// DeliveryListOrders contains list of orders for which delivery notes should be exported
type DeliveryListOrders struct {
	Orders []int `json: "orders"`
}

func (fapi *FyndiqAPI) deliveryNotesURL(params RequestParams) string {
	return fapi.getURL(
		getPath([]string{orderSegment}),
		params,
		true,
	)
}

// GetDeliveryNotesPDF returns HTTP Response object containing PDF with delivery notes
// http://fyndiq.github.io/api-v1/#resource-delivery-notes
func (fapi *FyndiqAPI) GetDeliveryNotesPDF(params RequestParams, orders *DeliveryListOrders) (*http.Response, error) {
	post, err := json.Marshal(orders)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", fapi.deliveryNotesURL(params), bytes.NewBuffer(post))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return resp, err
}
