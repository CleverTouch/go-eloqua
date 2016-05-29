package eloqua

import (
	"fmt"
)

// EmailService provides access to all the endpoints related
// to email assets within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Emails/emails-API.htm
type EmailService struct {
	client *Client
}

// Email represents an Eloqua email object.
type Email struct {
	AssetType       string   `json:"assetType,omitempty"`
	CurrentStatus   string   `json:"currentStatus,omitempty"`
	ID              int      `json:"id,omitempty,string"`
	CreatedAt       int      `json:"createdAt,omitempty,string"`
	CreatedBy       int      `json:"createdBy,omitempty,string"`
	RequestDepth    string   `json:"depth,omitempty"`
	FolderID        int      `json:"folderId,omitempty,string"`
	Name            string   `json:"name,omitempty"`
	Permissions     []string `json:"permissions,omitempty"`
	UpdatedAt       int      `json:"updatedAt,omitempty,string"`
	UpdatedBy       int      `json:"updatedBy,omitempty,string"`
	BounceBackEmail string   `json:"bouceBackEmail,omitempty"`
	// TODO - contentSections
	// TODO - dynamicContents
	EmailFooterID int `json:"emailFooterId,omitempty,string"`
	EmailHeaderID int `json:"emailHeaderId,omitempty,string"`
	EmailGroupID  int `json:"emailGroupId,omitempty,string"`
	EncodingID    int `json:"encodingId,omitempty,string"`
	// TODO - forms
	FieldMerges []FieldMerge `json:"fieldMerges,omitempty"`
	HTMLContent HTMLContent  `json:"htmlContent,omitempty"`
	Hyperlinks  []Hyperlink  `json:"hyperlinks,omitempty"`
	Images      []Image      `json:"images,omitempty"`

	PlainTextEditable bool   `json:"isPlainTextEditable,omitempty",string`
	Tracked           bool   `json:"isTracked,omitempty,string"`
	Subject           string `json:"subject,omitempty"`
	// TODO - landing pages
	Layout            string `json:"layout,omitempty"`
	PlainText         string `json:"plainText,omitempty"`
	ReplyToEmail      string `json:"replyToEmail,omitempty"`
	ReplyToName       string `json:"replyToName,omitempty"`
	SendPlainTextOnly bool   `json:"sendPlainTextOnly,omitempty,string"`
	SenderEmail       string `json:"senderEmail,omitempty"`
	SenderName        string `json:"senderName,omitempty"`
	Style             string `json:"style,omitempty"`
}

// HTMLContent represents the htmlContent model of an Eloqua email object
type HTMLContent struct {
	ContentType string `json:"type,omitempty"`
	HTML        string `json:"html,omitempty"`
}

// Create a new email in eloqua
func (e *EmailService) Create(name string, email *Email) (*Email, *Response, error) {
	if email == nil {
		email = &Email{}
	}
	email.Name = name
	endpoint := "/assets/email"
	resp, err := e.client.postRequestDecode(endpoint, email)
	return email, resp, err
}

// Get an email object via its ID
func (e *EmailService) Get(id int) (*Email, *Response, error) {
	endpoint := fmt.Sprintf("/assets/email/%d?depth=complete", id)
	email := &Email{}
	resp, err := e.client.getRequestDecode(endpoint, email)
	return email, resp, err
}

// List many Eloqua email objetcs
func (e *EmailService) List(opts *ListOptions) ([]Email, *Response, error) {
	endpoint := "/assets/emails"
	emails := new([]Email)
	resp, err := e.client.getRequestListDecode(endpoint, emails, opts)
	return *emails, resp, err
}

// Update an existing email in eloqua
func (e *EmailService) Update(id int, name string, email *Email) (*Email, *Response, error) {
	if email == nil {
		email = &Email{}
	}
	email.ID = id
	email.Name = name
	endpoint := fmt.Sprintf("/assets/email/%d", email.ID)
	resp, err := e.client.putRequestDecode(endpoint, email)
	return email, resp, err
}

// Delete an existing email from eloqua
func (e *EmailService) Delete(id int) (*Response, error) {
	email := &Email{ID: id}
	endpoint := fmt.Sprintf("/assets/email/%d", email.ID)
	resp, err := e.client.deleteRequest(endpoint, email)
	return resp, err
}
