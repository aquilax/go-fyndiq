package fyndiq

import (
	"encoding/json"
	"strconv"
)

const (
	PRODUCT_SEGMENT = "product"
)

type Variation struct {
	Id             int    `json: "id"`
	Name           string `json: "name"`
	NumInStock     int    `json: "num_in_stock"`
	Location       string `json: "location"`
	ItemNo         string `json: "item_no"`
	PlatformItemNo string `json: "platform_item_no"`
}

type ArticleGroup struct {
	Name       string      `json: "name"`
	Variations []Variation `json: "variations"`
}

type Product struct {
	Title             string       `json: "title"`
	Description       string       `json: "description"`
	Oldprice          string       `json: "oldprice"`
	Price             string       `json: "price"`
	MomsPercent       int          `json: "moms_percent"`
	NumInStock        int          `json: "num_in_stock"`
	State             string       `json: "state"`
	IsBlockedByFyndiq bool         `json: "is_blocked_by_fyndiq"`
	ItemNo            string       `json: "item_no"`
	PlatformItemNo    string       `json: "platform_item_no"`
	Location          string       `json: "location"`
	Url               string       `json: "url"`
	VariationGroup    ArticleGroup `json: "variation_group"`
	Images            []string     `json: "images"`
}

type ProductList struct {
	MetaResponse
	Objects []Product `json: "objects"`
}

// http://fyndiq.github.io/api-v1/#get-read-products
func (fapi *FyndiqApi) GetProducts() (*ProductList, error) {
	url := fapi.getUrl(
		getPath([]string{PRODUCT_SEGMENT}),
		RequestParams{},
	)
	jsonBlob, err := getJsonBlob(url)
	if err != nil {
		return nil, err
	}
	var result *ProductList
	return result, json.Unmarshal(jsonBlob, &result)
}

// http://fyndiq.github.io/api-v1/#get-read-products
func (fapi *FyndiqApi) GetProduct(id int) (*Product, error) {
	url := fapi.getUrl(
		getPath([]string{PRODUCT_SEGMENT, strconv.Itoa(id)}),
		RequestParams{},
	)
	jsonBlob, err := getJsonBlob(url)
	if err != nil {
		return nil, err
	}
	var result *Product
	return result, json.Unmarshal(jsonBlob, &result)

}
