package fyndiq

import (
	"bytes"
	"encoding/json"
	"strconv"
)

const (
	orderSegment = "order"
)

// OrderRow contains single order detail line
type OrderRow struct {
	Title                string  `json: "title"`
	Quantity             int     `json: "quantity"`
	ItemNo               string  `json: "item_no"`
	Location             string  `json: "location"`
	ProductID            int     `json: "product_id"`
	PricePerItemInclVat  float32 `json: "price_per_item_incl_vat"`
	PricePerItemExclVat  float32 `json: "price_per_item_excl_vat"`
	PricePerItemVat      float32 `json: "price_per_item_vat"`
	PriceSumInclVat      float32 `json: "price_sum_incl_vat"`
	PriceSumExclVat      float32 `json: "price_sum_excl_vat"`
	PriceSumVat          float32 `json: "price_sum_vat"`
	ProvisionInclVat     float32 `json: "provision_incl_vat"`
	ProvisionExclVat     float32 `json: "provision_excl_vat"`
	ProvisionVat         float32 `json: "provision_vat"`
	MerchantClaimInclVat float32 `json: "merchantclaim_incl_vat"`
	MerchantClaimExclVat float32 `json: "merchantclaim_excl_vat"`
	MerchantClaimVat     float32 `json: "merchantclaim_vat"`
	MomsPercent          int     `json: "moms_percent"`
	PlatformItemNo       string  `json: "platform_item_no"`
}

// Order contains the order header attributes
type Order struct {
	ID                   int        `json: "id"`
	InvoiceFirstName     string     `json: "invoice_firstname"`
	InvoiceLastName      string     `json: "invoice_lastname"`
	InvoiceAddress       string     `json: "invoice_address"`
	InvoicePostalCode    string     `json: "invoice_postalcode"`
	InvoiceCity          string     `json: "invoice_city"`
	InvoicePhone         string     `json: "invoice_phone"`
	InvoiceEmail         string     `json: "invoice_email"`
	InvoiceCO            string     `json: "invoice_co"`
	DeliveryFirstName    string     `json: "delivery_firstname"`
	DeliveryLastName     string     `json: "delivery_lastname"`
	DeliveryAddress      string     `json: "delivery_address"`
	DeliveryPostalCode   string     `json: "delivery_postalcode"`
	DeliveryCity         string     `json: "delivery_city"`
	DeliveryCO           string     `json: "delivery_co"`
	PdfURL               string     `json: "pdf_url"`
	FreightbackInclVAT   float32    `json: "freightback_incl_vat"`
	FreightbackExclVAT   float32    `json: "freightback_excl_vat"`
	FreightbackVAT       float32    `json: "freightback_vat"`
	ProvisionInclVAT     float32    `json: "provision_incl_vat"`
	ProvisionEcclVAT     float32    `json: "provision_excl_vat"`
	ProvisionVAT         float32    `json: "provision_vat"`
	MerchantClaimInclVAT float32    `json: "merchantclaim_incl_vat"`
	MerchantClaimExclVAT float32    `json: "merchantclaim_excl_vat"`
	Marked               bool       `json: "marked"`
	OrderRows            []OrderRow `json: "order_rows"`
}

// OrderList represents list of orders with meta data
type OrderList struct {
	MetaResponse
	Objects []Order `json: "objects"`
}

// GetOrders fetches list of all orders
// http://fyndiq.github.io/api-v1/#resource-order
func (fapi *FyndiqAPI) GetOrders(params RequestParams) (*OrderList, error) {
	url := fapi.getURL(
		getPath([]string{orderSegment}),
		params,
	)
	resp, err := httpRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result *OrderList
	return result, json.Unmarshal(resp.body, &result)
}

// UpdateOrder updates existing order
// http://fyndiq.github.io/api-v1/#resource-order
func (fapi *FyndiqAPI) UpdatOrder(id int, updateData []byte) error {
	url := fapi.getURL(
		getPath([]string{orderSegment, strconv.Itoa(id)}),
		RequestParams{},
	)
	_, err := httpRequest("PUT", url, bytes.NewBuffer(updateData))
	return err
}
