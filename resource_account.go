package fyndiq

import (
	"bytes"
	"encoding/json"
	"strconv"
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

func (fapi *FyndiqAPI) getAccountURL() string {
	return fapi.getURL(
		getPath([]string{accountSegment}),
		RequestParams{},
		true,
	)
}

func (fapi *FyndiqAPI) updateAccountURL(id int) string {
	return fapi.getURL(
		getPath([]string{accountSegment, strconv.Itoa(id)}),
		RequestParams{},
		true,
	)
}

func (fapi *FyndiqAPI) createAccountURL() string {
	return fapi.getURL(
		getPath([]string{accountSegment}),
		RequestParams{},
		false,
	)
}

func (fapi *FyndiqAPI) GetAccount() (*Account, error) {
	resp, err := httpRequest("GET", fapi.getAccountURL(), nil)
	if err != nil {
		return nil, err
	}
	var result *Account
	return result, json.Unmarshal(resp.body, &result)
}

func (fapi *FyndiqAPI) UpdateAccount(id int, updateData []byte) error {
	_, err := httpRequest("PUT", fapi.updateAccountURL(id), bytes.NewBuffer(updateData))
	return err
}

func (fapi *FyndiqAPI) CreateAccount(account *Account) (string, error) {
	post, err := json.Marshal(account)
	if err != nil {
		return "", err
	}
	resp, err := httpRequest("POST", fapi.createAccountURL(), bytes.NewBuffer(post))
	if err == nil {
		return resp.header["Location"][0], nil
	}
	return "", err
}
