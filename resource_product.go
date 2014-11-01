package fyndiq

import (
	"bytes"
	"encoding/json"
	"strconv"
)

const (
	productSegment = "product"
)

// Variation for article group
// http://fyndiq.github.io/api-v1/#product-article-group
type Variation struct {
	ID             int    `json: "id"`
	Name           string `json: "name"`
	NumInStock     int    `json: "num_in_stock"`
	Location       string `json: "location"`
	ItemNo         string `json: "item_no"`
	PlatformItemNo string `json: "platform_item_no"`
}

// ArticleGroup for product
// http://fyndiq.github.io/api-v1/#product-article
type ArticleGroup struct {
	Name       string      `json: "name"`
	Variations []Variation `json: "variations"`
}

// Product represents single product
// http://fyndiq.github.io/api-v1/#product
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
	URL               string       `json: "url"`
	VariationGroup    ArticleGroup `json: "variation_group"`
	Images            []string     `json: "images"`
}

// ProductList represents list of products with meta data
type ProductList struct {
	MetaResponse
	Objects []Product `json: "objects"`
}

// GetProducts fetches list of all products
// http://fyndiq.github.io/api-v1/#get-read-products
func (fapi *FyndiqAPI) GetProducts() (*ProductList, error) {
	url := fapi.getURL(
		getPath([]string{productSegment}),
		RequestParams{},
	)
	resp, err := httpRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result *ProductList
	return result, json.Unmarshal(resp.body, &result)
}

// GetProduct fetches single product by ID
// http://fyndiq.github.io/api-v1/#get-read-products
func (fapi *FyndiqAPI) GetProduct(id int) (*Product, error) {
	url := fapi.getURL(
		getPath([]string{productSegment, strconv.Itoa(id)}),
		RequestParams{},
	)
	resp, err := httpRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result *Product
	return result, json.Unmarshal(resp.body, &result)
}

// DeleteProduct deletes single product by ID
// http://fyndiq.github.io/api-v1/#delete-delete-products
func (fapi *FyndiqAPI) DeleteProduct(id int) error {
	url := fapi.getURL(
		getPath([]string{productSegment, strconv.Itoa(id)}),
		RequestParams{},
	)
	_, err := httpRequest("DELETE", url, nil)
	return err
}

// CreateProduct creates new product
// http://fyndiq.github.io/api-v1/#post-create-products
func (fapi *FyndiqAPI) CreateProduct(product *Product) (string, error) {
	url := fapi.getURL(
		getPath([]string{productSegment}),
		RequestParams{},
	)
	post, err := json.Marshal(product)
	if err != nil {
		return "", err
	}
	resp, err := httpRequest("POST", url, bytes.NewBuffer(post))
	if err == nil {
		return resp.header["Location"][0], nil
	}
	return "", err
}
