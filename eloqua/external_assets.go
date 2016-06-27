package eloqua

import (
	"fmt"
)

// ExternalAssetService provides access to all the endpoints related
// to External Assets within eloqua
//
// Eloqua API docs: https://goo.gl/lntJIr
type ExternalAssetService struct {
	client *Client
}

// ExternalAsset represents an Eloqua ExternalAsset object.
type ExternalAsset struct {
	Type                string `json:"type,omitempty"`
	ID                  int    `json:"id,omitempty,string"`
	CreatedAt           int    `json:"createdAt,omitempty,string"`
	CreatedBy           int    `json:"createdBy,omitempty,string"`
	Depth               string `json:"depth,omitempty"`
	Name                string `json:"name,omitempty"`
	UpdatedAt           int    `json:"updatedAt,omitempty,string"`
	UpdatedBy           int    `json:"updatedBy,omitempty,string"`
	ExternalAssetTypeID int    `json:"externalAssetTypeId,omitempty,string"`
}

// Create a new externalAsset in eloqua
func (e *ExternalAssetService) Create(name string, externalAsset *ExternalAsset) (*ExternalAsset, *Response, error) {
	if externalAsset == nil {
		externalAsset = &ExternalAsset{}
	}
	externalAsset.Name = name

	endpoint := "/assets/external"
	resp, err := e.client.postRequestDecode(endpoint, externalAsset)
	return externalAsset, resp, err
}

// Get an externalAsset object via its ID
func (e *ExternalAssetService) Get(id int) (*ExternalAsset, *Response, error) {
	endpoint := fmt.Sprintf("/assets/external/%d?depth=complete", id)
	externalAsset := &ExternalAsset{}
	resp, err := e.client.getRequestDecode(endpoint, externalAsset)
	return externalAsset, resp, err
}

// List many eloqua externalAssets
func (e *ExternalAssetService) List(opts *ListOptions) ([]ExternalAsset, *Response, error) {
	endpoint := "/assets/externals"
	externalAssets := new([]ExternalAsset)
	resp, err := e.client.getRequestListDecode(endpoint, externalAssets, opts)
	return *externalAssets, resp, err
}

// Update an existing externalAsset in eloqua
func (e *ExternalAssetService) Update(id int, name string, externalAsset *ExternalAsset) (*ExternalAsset, *Response, error) {
	if externalAsset == nil {
		externalAsset = &ExternalAsset{}
	}

	externalAsset.ID = id
	externalAsset.Name = name

	endpoint := fmt.Sprintf("/assets/external/%d", externalAsset.ID)
	resp, err := e.client.putRequestDecode(endpoint, externalAsset)
	return externalAsset, resp, err
}

// Delete an existing externalAsset from eloqua
// During testing this did not seem to function but it is
// in the documentation and does not return an error so it will remain for now.
func (e *ExternalAssetService) Delete(id int) (*Response, error) {
	externalAsset := &ExternalAsset{ID: id}
	endpoint := fmt.Sprintf("/assets/external/%d", externalAsset.ID)
	resp, err := e.client.deleteRequest(endpoint, externalAsset)
	return resp, err
}
