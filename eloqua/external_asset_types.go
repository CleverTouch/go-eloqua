package eloqua

import (
	"fmt"
)

// ExternalAssetTypeService provides access to all the endpoints related
// to External Assets within eloqua
//
// Eloqua API docs: https://goo.gl/UhnlB0
type ExternalAssetTypeService struct {
	client *Client
}

// ExternalAssetType represents an Eloqua ExternalAssetType object.
type ExternalAssetType struct {
	Type          string                 `json:"type,omitempty"`
	ID            int                    `json:"id,omitempty,string"`
	CreatedAt     int                    `json:"createdAt,omitempty,string"`
	CreatedBy     int                    `json:"createdBy,omitempty,string"`
	Depth         string                 `json:"depth,omitempty"`
	Name          string                 `json:"name,omitempty"`
	UpdatedAt     int                    `json:"updatedAt,omitempty,string"`
	UpdatedBy     int                    `json:"updatedBy,omitempty,string"`
	ActivityTypes []ExternalActivityType `json:"activityTypes,omitempty,string"`
}

// Create a new externalAssetType in eloqua.
// New activity types can be created by sending them through this request.
func (e *ExternalAssetTypeService) Create(name string, externalAssetType *ExternalAssetType) (*ExternalAssetType, *Response, error) {
	if externalAssetType == nil {
		externalAssetType = &ExternalAssetType{}
	}
	externalAssetType.Name = name

	endpoint := "/assets/external/type"
	resp, err := e.client.postRequestDecode(endpoint, externalAssetType)
	return externalAssetType, resp, err
}

// Get an externalAssetType object via its ID
func (e *ExternalAssetTypeService) Get(id int) (*ExternalAssetType, *Response, error) {
	endpoint := fmt.Sprintf("/assets/external/type/%d?depth=complete", id)
	externalAssetType := &ExternalAssetType{}
	resp, err := e.client.getRequestDecode(endpoint, externalAssetType)
	return externalAssetType, resp, err
}

// List many eloqua externalAssetTypes
func (e *ExternalAssetTypeService) List(opts *ListOptions) ([]ExternalAssetType, *Response, error) {
	endpoint := "/assets/external/types"
	externalAssetTypes := new([]ExternalAssetType)
	resp, err := e.client.getRequestListDecode(endpoint, externalAssetTypes, opts)
	return *externalAssetTypes, resp, err
}

// Update an existing externalAssetType in eloqua.
// New activity types can be created by sending them through this request.
func (e *ExternalAssetTypeService) Update(id int, name string, externalAssetType *ExternalAssetType) (*ExternalAssetType, *Response, error) {
	if externalAssetType == nil {
		externalAssetType = &ExternalAssetType{}
	}

	externalAssetType.ID = id
	externalAssetType.Name = name

	endpoint := fmt.Sprintf("/assets/external/type/%d", externalAssetType.ID)
	resp, err := e.client.putRequestDecode(endpoint, externalAssetType)
	return externalAssetType, resp, err
}

// Delete an existing externalAssetType from eloqua
func (e *ExternalAssetTypeService) Delete(id int) (*Response, error) {
	externalAssetType := &ExternalAssetType{ID: id}
	endpoint := fmt.Sprintf("/assets/external/type/%d", externalAssetType.ID)
	resp, err := e.client.deleteRequest(endpoint, externalAssetType)
	return resp, err
}
