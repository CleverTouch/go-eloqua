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
	FolderId        int      `json:"folderId,omitempty,string"`
	Name            string   `json:"name,omitempty"`
	Permissions     []string `json:"permissions,omitempty"`
	UpdatedAt       int      `json:"updatedAt,omitempty,string"`
	BounceBackEmail string   `json:"bouceBackEmail,omitempty"`
	// TODO - contentSections
	// TODO - dynamicContents
	EmailFooterId int `json:"emailFooterId,omitempty,string"`
	EmailHeaderId int `json:"emailHeaderId,omitempty,string"`
	EmailGroupId  int `json:"emailGroupId,omitempty,string"`
	EncodingId    int `json:"encodingId,omitempty,string"`
	// TODO - fieldMerges
	// TODO - forms
	HtmlContent htmlContent `json:"htmlContent,omitempty"`
	// TODO - hyperlinks
	// TODO - images
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

// htmlContent represents the htmlContent component of an email
type htmlContent struct {
	ContentType string `json:"assetType,omitempty"`
	Html        string `json:"html,omitempty"`
}

// Get an email object via its ID
func (e *EmailService) Get(id int) (*Email, *Response, error) {
	endpoint := fmt.Sprintf("/assets/email/%d?depth=complete", id)
	email := &Email{}
	resp, err := e.client.getRequestDecode(endpoint, email)
	return email, resp, err
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