package eloqua

import (
	"fmt"
)

// LandingPageService provides access to all the endpoints related
// to landingPage assets within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Landing pages/landingPages-API.htm
type LandingPageService struct {
	client *Client
}

// LandingPage represents an Eloqua landingPage object.
type LandingPage struct {
	Type                string           `json:"type,omitempty"`
	CurrentStatus       string           `json:"currentStatus,omitempty"`
	ID                  int              `json:"id,omitempty,string"`
	CreatedAt           int              `json:"createdAt,omitempty,string"`
	CreatedBy           int              `json:"createdBy,omitempty,string"`
	Depth               string           `json:"depth,omitempty"`
	FolderID            int              `json:"folderId,omitempty,string"`
	Name                string           `json:"name,omitempty"`
	Permissions         []string         `json:"permissions,omitempty"`
	UpdatedAt           int              `json:"updatedAt,omitempty,string"`
	UpdatedBy           int              `json:"updatedBy,omitempty,string"`
	AutoRedirectURL     string           `json:"autoRedirectURL,omitempty"`
	AutoRedirectWaitFor int              `json:"autoRedirectWaitFor,omitempty,string"`
	ContentSections     []ContentSection `json:"contentSections,omitempty"`
	DeployedAt          int              `json:"deployedAt,omitempty,string"`
	DynamicContents     []DynamicContent `json:"dynamicContents,omitempty"`
	Forms               []Form           `json:"forms,omitempty"`
	HTMLContent         HTMLContent      `json:"htmlContent,omitempty"`
	Hyperlinks          []Hyperlink      `json:"hyperlinks,omitempty"`
	Images              []Image          `json:"images,omitempty"`
	Layout              string           `json:"layout,omitempty"`
	MicrositeId         int              `json:"micrositeId,omitempty,string"`
	Style               string           `json:"style,omitempty"`
	RefreshedAt         int              `json:"refreshedAt,omitempty,string"`
	RelativePath        string           `json:"relativePath,omitempty"`

	IsContentProtected        bool `json:"isContentProtected,omitempty,string"`
	ExcludeFromAuthentication bool `json:"excludeFromAuthentication,omitempty,string"`
}

// Create a new landingPage in eloqua
func (e *LandingPageService) Create(name string, landingPage *LandingPage) (*LandingPage, *Response, error) {
	if landingPage == nil {
		landingPage = &LandingPage{}
	}
	landingPage.Name = name
	endpoint := "/assets/landingPage"
	resp, err := e.client.postRequestDecode(endpoint, landingPage)
	return landingPage, resp, err
}

// Get an landingPage object via its ID
func (e *LandingPageService) Get(id int) (*LandingPage, *Response, error) {
	endpoint := fmt.Sprintf("/assets/landingPage/%d?depth=complete", id)
	landingPage := &LandingPage{}
	resp, err := e.client.getRequestDecode(endpoint, landingPage)
	return landingPage, resp, err
}

// List many Eloqua landingPage objetcs
func (e *LandingPageService) List(opts *ListOptions) ([]LandingPage, *Response, error) {
	endpoint := "/assets/landingPages"
	landingPages := new([]LandingPage)
	resp, err := e.client.getRequestListDecode(endpoint, landingPages, opts)
	return *landingPages, resp, err
}

// Update an existing landingPage in eloqua
func (e *LandingPageService) Update(id int, name string, landingPage *LandingPage) (*LandingPage, *Response, error) {
	if landingPage == nil {
		landingPage = &LandingPage{}
	}
	landingPage.ID = id
	landingPage.Name = name
	endpoint := fmt.Sprintf("/assets/landingPage/%d", landingPage.ID)
	resp, err := e.client.putRequestDecode(endpoint, landingPage)
	return landingPage, resp, err
}

// Delete an existing landingPage from eloqua
func (e *LandingPageService) Delete(id int) (*Response, error) {
	landingPage := &LandingPage{ID: id}
	endpoint := fmt.Sprintf("/assets/landingPage/%d", landingPage.ID)
	resp, err := e.client.deleteRequest(endpoint, landingPage)
	return resp, err
}
