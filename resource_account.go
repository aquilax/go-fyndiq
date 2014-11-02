package fyndiq

import (
	"encoding/json"
)

const (
	accountSegment = "account"
)

type Account struct {
	UserName       string `json: "username"`
	Password       string `json: "password"`
	UserFirstName  string `json: "user_firstname"`
	UserLastName   string `json: "user_lastname"`
	UserEmail      string `json: "user_email"`
	UserPhone      string `json: "user_phone"`
	ShopName       string `json: "shop_name"`
	ShopCompany    string `json: "shop_company"`
	ShopUrl        string `json: "shop_url"`
	ShopOrgNum     string `json: "shop_org_num"`
	ShopVatNum     string `json: "shop_vat_num"`
	ShopAddress    string `json: "shop_address"`
	ShopPostalCode string `json: "shop_postal_code"`
	ShopCity       string `json: "shop_city"`
	ShopPhone      string `json: "shop_phone"`
	ID             int    `json: "id"`
	NotifyURL      string `json: "notify_url"`
	NotifyAnswer   string `json: "notify_answer"`
}

func (fapi *FyndiqAPI) GetAccount() (*Account, error) {
	url := fapi.getURL(
		getPath([]string{accountSegment}),
		RequestParams{},
	)
	resp, err := httpRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result *Account
	return result, json.Unmarshal(resp.body, &result)
}
