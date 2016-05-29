package eloqua

import (
	"fmt"
)

// FormDataService provides access to all the endpoints related
// to form data within eloqua.
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Form data/formData-API.htm
type FormDataService struct {
	client *Client
}

// FormData represents an Eloqua form record.
type FormData struct {
	Type                 string       `json:"type,omitempty"`
	Name                 string       `json:"name,omitempty"`
	ID                   int          `json:"id,omitempty,string"`
	FieldValues          []FieldValue `json:"fieldValues,omitempty"`
	SubmittedAt          int          `json:"submittedAt,omitempty,string"`
	SubmittedByContactID int          `json:"submittedByContactId,omitempty,string"`
}

// Create a new form record in eloqua
func (e *FormDataService) Create(formID int, formData *FormData) (*FormData, *Response, error) {
	if formData == nil {
		formData = &FormData{}
	}

	endpoint := fmt.Sprintf("/data/form/%d", formID)
	resp, err := e.client.postRequestDecode(endpoint, formData)
	return formData, resp, err
}

// List many eloqua form records
func (e *FormDataService) List(formID int, opts *ListOptions) ([]FormData, *Response, error) {
	endpoint := fmt.Sprintf("/data/form/%d", formID)
	formDatas := new([]FormData)
	resp, err := e.client.getRequestListDecode(endpoint, formDatas, opts)
	return *formDatas, resp, err
}
