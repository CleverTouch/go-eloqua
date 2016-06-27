package eloqua

import (
	"fmt"
)

// CustomObjectDataService provides access to all the endpoints related
// to custom object data within eloqua
//
// Eloqua API docs: https://goo.gl/XmjLBw
type CustomObjectDataService struct {
	client *Client
}

// CustomObjectData represents an Eloqua custom object record.
// Some of the endpoints for CustomObjectData are not official listed
// within the Eloqua documentation so may be more prone to breaking.
type CustomObjectData struct {
	Type        string       `json:"type,omitempty"`
	Name        string       `json:"name,omitempty"`
	ID          int          `json:"id,omitempty,string"`
	FieldValues []FieldValue `json:"fieldValues,omitempty"`
	UniqueCode  string       `json:"uniqueCode,omitempty"`
	CreatedAt   int          `json:"createdAt,omitempty,string"`
}

// Create a new custom object record in eloqua
func (e *CustomObjectDataService) Create(cdoID int, customObjectData *CustomObjectData) (*CustomObjectData, *Response, error) {
	if customObjectData == nil {
		customObjectData = &CustomObjectData{}
	}

	endpoint := fmt.Sprintf("/data/customObject/%d/instance", cdoID)
	resp, err := e.client.postRequestDecode(endpoint, customObjectData)
	return customObjectData, resp, err
}

// Get a custom object data record via its ID, Within the CDO of the given cdoID.
func (e *CustomObjectDataService) Get(cdoID int, id int) (*CustomObjectData, *Response, error) {
	endpoint := fmt.Sprintf("/data/customObject/%d/instance/%d?depth=complete", cdoID, id)
	customObjectData := &CustomObjectData{}
	resp, err := e.client.getRequestDecode(endpoint, customObjectData)
	return customObjectData, resp, err
}

// List many eloqua custom object records
func (e *CustomObjectDataService) List(cdoID int, opts *ListOptions) ([]CustomObjectData, *Response, error) {
	endpoint := fmt.Sprintf("/data/customObject/%d/instances", cdoID)
	customObjectDatas := new([]CustomObjectData)
	resp, err := e.client.getRequestListDecode(endpoint, customObjectDatas, opts)
	return *customObjectDatas, resp, err
}

// Update an existing custom object in eloqua
// To actually update the cdo record value ensure you pass a customObjectData model
// with its FieldValues filled.
func (e *CustomObjectDataService) Update(cdoID int, id int, customObjectData *CustomObjectData) (*CustomObjectData, *Response, error) {
	if customObjectData == nil {
		customObjectData = &CustomObjectData{}
	}

	customObjectData.ID = id

	endpoint := fmt.Sprintf("/data/customObject/%d/instance/%d", cdoID, customObjectData.ID)
	resp, err := e.client.putRequestDecode(endpoint, customObjectData)
	return customObjectData, resp, err
}

// Delete an existing custom object record from eloqua
func (e *CustomObjectDataService) Delete(cdoID int, id int) (*Response, error) {
	customObjectData := &CustomObjectData{ID: id}
	endpoint := fmt.Sprintf("/data/customObject/%d/instance/%d", cdoID, customObjectData.ID)
	resp, err := e.client.deleteRequest(endpoint, customObjectData)
	return resp, err
}
