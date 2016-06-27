package eloqua

import (
	"fmt"
)

// AccountService provides access to all the endpoints related
// to account data within eloqua
//
// Eloqua API docs: https://goo.gl/s1bDwX
type AccountService struct {
	client *Client
}

// Account represents an Eloqua email object.
// Fields that are not listed in the Account model itself can be retrieved/updated
// using the 'FieldValues' property.
type Account struct {
	Type          string `json:"type,omitempty"`
	CurrentStatus string `json:"currentStatus,omitempty"`
	ID            int    `json:"id,omitempty,string"`
	CreatedAt     int    `json:"createdAt,omitempty,string"`
	CreatedBy     int    `json:"createdBy,omitempty,string"`
	Depth         string `json:"depth,omitempty"`
	UpdatedAt     int    `json:"updatedAt,omitempty,string"`
	UpdatedBy     int    `json:"updatedBy,omitempty,string"`

	Name          string `json:"name,omitempty"`
	Address1      string `json:"address1,omitempty"`
	Address2      string `json:"address2,omitempty"`
	Address3      string `json:"address3,omitempty"`
	BusinessPhone string `json:"businessPhone,omitempty"`
	City          string `json:"city,omitempty"`
	Country       string `json:"country,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
	Province      string `json:"province,omitempty"`

	FieldValues []FieldValue `json:"fieldValues,imitempty"`
}

// Create a new account in eloqua
func (e *AccountService) Create(name string, account *Account) (*Account, *Response, error) {
	if account == nil {
		account = &Account{}
	}
	account.Name = name
	endpoint := "/data/account"
	resp, err := e.client.postRequestDecode(endpoint, account)
	return account, resp, err
}

// Get an account object via its ID
func (e *AccountService) Get(id int) (*Account, *Response, error) {
	endpoint := fmt.Sprintf("/data/account/%d?depth=complete", id)
	account := &Account{}
	resp, err := e.client.getRequestDecode(endpoint, account)
	return account, resp, err
}

// List many Eloqua account objects
func (e *AccountService) List(opts *ListOptions) ([]Account, *Response, error) {
	endpoint := "/data/accounts"
	accounts := new([]Account)
	resp, err := e.client.getRequestListDecode(endpoint, accounts, opts)
	return *accounts, resp, err
}

// Update an existing account in eloqua
func (e *AccountService) Update(id int, name string, account *Account) (*Account, *Response, error) {
	if account == nil {
		account = &Account{}
	}
	account.ID = id
	account.Name = name
	endpoint := fmt.Sprintf("/data/account/%d", account.ID)
	resp, err := e.client.putRequestDecode(endpoint, account)
	return account, resp, err
}

// Delete an existing account from eloqua
func (e *AccountService) Delete(id int) (*Response, error) {
	account := &Account{ID: id}
	endpoint := fmt.Sprintf("/data/account/%d", account.ID)
	resp, err := e.client.deleteRequest(endpoint, account)
	return resp, err
}
