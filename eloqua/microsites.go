package eloqua

import (
	"fmt"
)

// MicrositeService provides access to all the endpoints related
// to microsite data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Microsites/microsites-API.htm
type MicrositeService struct {
	client *Client
}

// Microsite represents an Eloqua microsite object.
// Microsites are often in other Eloqua models such as Emails & Landing Pages
type Microsite struct {
	Type      string `json:"type,omitempty"`
	ID        int    `json:"id,omitempty,string"`
	CreatedAt int    `json:"createdAt,omitempty,string"`
	CreatedBy int    `json:"createdBy,omitempty,string"`
	Depth     string `json:"depth,omitempty"`
	Name      string `json:"name,omitempty"`
	UpdatedAt int    `json:"updatedAt,omitempty,string"`

	Domains                []string `json:"domains,omitempty"`
	EnableWebTrackingOptIn string   `json:"enableWebTrackingOptIn,omitempty"`
	IsAuthenticated        bool     `json:"isAuthenticated,omitempty,string"`
	IsSecure               bool     `json:"isSecure,omitempty,string"`
}

// Create a new microsite in eloqua
func (e *MicrositeService) Create(name string, microsite *Microsite) (*Microsite, *Response, error) {
	if microsite == nil {
		microsite = &Microsite{}
	}
	microsite.Name = name

	endpoint := "/assets/microsite"
	resp, err := e.client.postRequestDecode(endpoint, microsite)
	return microsite, resp, err
}

// Get an microsite object via its ID
func (e *MicrositeService) Get(id int) (*Microsite, *Response, error) {
	endpoint := fmt.Sprintf("/assets/microsite/%d?depth=complete", id)
	microsite := &Microsite{}
	resp, err := e.client.getRequestDecode(endpoint, microsite)
	return microsite, resp, err
}

// List many eloqua microsites
func (e *MicrositeService) List(opts *ListOptions) ([]Microsite, *Response, error) {
	endpoint := "/assets/microsites"
	microsites := new([]Microsite)
	resp, err := e.client.getRequestListDecode(endpoint, microsites, opts)
	return *microsites, resp, err
}

// Update an existing microsite in eloqua
func (e *MicrositeService) Update(id int, name string, microsite *Microsite) (*Microsite, *Response, error) {
	if microsite == nil {
		microsite = &Microsite{}
	}

	microsite.ID = id
	microsite.Name = name

	endpoint := fmt.Sprintf("/assets/microsite/%d", microsite.ID)
	resp, err := e.client.putRequestDecode(endpoint, microsite)
	return microsite, resp, err
}

// Delete an existing microsite from eloqua
func (e *MicrositeService) Delete(id int) (*Response, error) {
	microsite := &Microsite{ID: id}
	endpoint := fmt.Sprintf("/assets/microsite/%d", microsite.ID)
	resp, err := e.client.deleteRequest(endpoint, microsite)
	return resp, err
}
