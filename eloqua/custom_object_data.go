package eloqua

import (
	"fmt"
)

// CustomObjectDataService provides access to all the endpoints related
// to custom object data  within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Custom objects/customObjectDataDatas-API.htm
type CustomObjectDataService struct {
	client *Client
}

// CustomObjectDataList is the returned list from Eloqua when looking up the
// data on a custom data object
type CustomObjectDataList struct {
	Elements []CustomObjectData `json:"elements,omitempty"`
	Page     int                `json:"page,omitempty"`
	PageSize int                `json:"pageSize,omitempty"`
	Total    int                `json:"total,omitempty"`
}

// CustomObjectData represents an Eloqua custom object.
type CustomObjectData struct {
	Type        string       `json:"type,omitempty"`
	ID          int          `json:"id,omitempty,string"`
	FieldValues []FieldValue `json:"fieldValues,omitempty"`
}

// Create a new custom object in eloqua
func (e *CustomObjectDataService) Create(name string, customObjectData *CustomObjectData) (*CustomObjectData, *Response, error) {
	if customObjectData == nil {
		customObjectData = &CustomObjectData{}
	}

	customObjectData.Name = name
	endpoint := "/assets/customObjectData"
	resp, err := e.client.postRequestDecode(endpoint, customObjectData)
	return customObjectData, resp, err
}

// Get a custom object via its ID
func (e *CustomObjectDataService) Get(id int) (*CustomObjectData, *Response, error) {
	endpoint := fmt.Sprintf("/assets/customObjectData/%d?depth=complete", id)
	customObjectData := &CustomObjectData{}
	resp, err := e.client.getRequestDecode(endpoint, customObjectData)
	return customObjectData, resp, err
}

// List many eloqua custom objects
func (e *CustomObjectDataService) List(opts *ListOptions) ([]CustomObjectData, *Response, error) {
	endpoint := "/assets/customObjectDatas"
	customObjectDatas := new([]CustomObjectData)
	resp, err := e.client.getRequestListDecode(endpoint, customObjectDatas, opts)
	return *customObjectDatas, resp, err
}

// Update an existing custom object in eloqua
func (e *CustomObjectDataService) Update(id int, name string, customObjectData *CustomObjectData) (*CustomObjectData, *Response, error) {
	if customObjectData == nil {
		customObjectData = &CustomObjectData{}
	}

	customObjectData.ID = id
	customObjectData.Name = name

	endpoint := fmt.Sprintf("/assets/customObjectData/%d", customObjectData.ID)
	resp, err := e.client.putRequestDecode(endpoint, customObjectData)
	return customObjectData, resp, err
}

// Delete an existing custom object from eloqua
func (e *CustomObjectDataService) Delete(id int) (*Response, error) {
	customObjectData := &CustomObjectData{ID: id}
	endpoint := fmt.Sprintf("/assets/customObjectData/%d", customObjectData.ID)
	resp, err := e.client.deleteRequest(endpoint, customObjectData)
	return resp, err
}
