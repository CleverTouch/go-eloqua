package eloqua

import (
	"fmt"
)

// ContactService provides access to all the endpoints related
// to contact data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Contacts/contacts-API.htm
type ContactService struct {
	client *Client
}

// Contact represents an Eloqua email object.
// Fields that are not listed in the Contact model itself can be retrieved/updated
// using the 'FieldValues' property.
type Contact struct {
	Type          string `json:"type,omitempty"`
	CurrentStatus string `json:"currentStatus,omitempty"`
	ID            int    `json:"id,omitempty,string"`
	CreatedAt     int    `json:"createdAt,omitempty,string"`
	RequestDepth  string `json:"depth,omitempty"`
	// This actually relates to the contact's email address
	// rather than the contacts name
	Name          string `json:"name,omitempty"`
	UpdatedAt     int    `json:"updatedAt,omitempty,string"`
	AccountName   string `json:"accountName,omitempty"`
	BusinessPhone string `json:"businessPhone,omitempty"`
	Country       string `json:"country,omitempty"`

	EmailAddress          string `json:"emailAddress,omitempty"`
	EmailFormatPreference string `json:"emailFormatPreference,omitempty"`
	FirstName             string `json:"firstName,omitempty"`
	LastName              string `json:"lastName,omitempty"`
	PostalCode            string `json:"postalCode,omitempty"`
	Province              string `json:"province,omitempty"`
	SalesPerson           string `json:"salesPerson,omitempty"`
	// Job title, Not name title
	Title string `json:"title,omitempty"`

	SubscriptionDate int          `json:"subscriptionDate,omitempty,string"`
	IsBounceBack     bool         `json:"isBounceBack,omitempty,string"`
	IsSubscribed     bool         `json:"isSubscribed,imitempty,string"`
	FieldValues      []FieldValue `json:"fieldValues,imitempty"`
}

// A FieldValue represents an Eloqua field values objects
type FieldValue struct {
	Type  string `json:"type,omitempty"`
	ID    int    `json:"id,omitempty,string"`
	Value string `json:"value,omitempty"`
}

// Create a new contact in eloqua
// The email must not already exists otherwise Eloqua will return an error.
func (e *ContactService) Create(emailAddress string, contact *Contact) (*Contact, *Response, error) {
	if contact == nil {
		contact = &Contact{}
	}
	contact.EmailAddress = emailAddress
	endpoint := "/data/contact"
	resp, err := e.client.postRequestDecode(endpoint, contact)
	return contact, resp, err
}

// Get an contact object via its ID
func (e *ContactService) Get(id int) (*Contact, *Response, error) {
	endpoint := fmt.Sprintf("/data/contact/%d?depth=complete", id)
	contact := &Contact{}
	resp, err := e.client.getRequestDecode(endpoint, contact)
	return contact, resp, err
}

// List many Eloqua contact objects
func (e *ContactService) List(opts *ListOptions) ([]Contact, *Response, error) {
	endpoint := "/data/contacts"
	contacts := new([]Contact)
	resp, err := e.client.getRequestListDecode(endpoint, contacts, opts)
	return *contacts, resp, err
}

// Update an existing contact in eloqua
func (e *ContactService) Update(id int, emailAddress string, contact *Contact) (*Contact, *Response, error) {
	if contact == nil {
		contact = &Contact{}
	}
	contact.ID = id
	contact.EmailAddress = emailAddress
	endpoint := fmt.Sprintf("/data/contact/%d", contact.ID)
	resp, err := e.client.putRequestDecode(endpoint, contact)
	return contact, resp, err
}

// Delete an existing contact from eloqua
func (e *ContactService) Delete(id int) (*Response, error) {
	contact := &Contact{ID: id}
	endpoint := fmt.Sprintf("/data/contact/%d", contact.ID)
	resp, err := e.client.deleteRequest(endpoint, contact)
	return resp, err
}
