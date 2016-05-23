package eloqua

import (
	"fmt"
)

// EmailGroupService provides access to all the endpoints related
// to email group data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Email groups/post-assets-emailGroup.htm
type EmailGroupService struct {
	client *Client
}

// EmailGroup represents an Eloqua email group object.
type EmailGroup struct {
	Type          string   `json:"type,omitempty"`
	ID            int      `json:"id,omitempty,string"`
	CreatedAt     int      `json:"createdAt,omitempty,string"`
	CreatedBy     int      `json:"createdBy,omitempty,string"`
	RequestDepth  string   `json:"depth,omitempty"`
	Name          string   `json:"name,omitempty"`
	Permissions   []string `json:"permissions,omitempty"`
	Description   string   `json:"description,omitempty"`
	UpdatedAt     int      `json:"updatedAt,omitempty,string"`
	UpdatedBy     int      `json:"updatedBy,omitempty,string"`
	EmailHeaderID int      `json:"emailHeaderId,omitempty,string"`
	EmailFooterID int      `json:"emailFooterId,omitempty,string"`
	EmailIDs      []int    `json:"emailIds,omitempty,string"`

	IsVisibleInOutlookPlugin          bool   `json:"isVisibleInOutlookPlugin,omitempty,string"`
	IsVisibleInPublicSubscriptionList bool   `json:"isVisibleInPublicSubscriptionList,omitempty,string"`
	SubscriptionListDataLookupID      string `json:"subscriptionListDataLookupId,omitempty"`
	SubscriptionListID                int    `json:"subscriptionListId,omitempty,string"`
	SubscriptionLandingPageId         int    `json:"subscriptionLandingPageId,omitempty,string"`
	UnSubscriptionListDataLookupId    string `json:"unSubscriptionListDataLookupId,omitempty"`
	UnSubscriptionListId              int    `json:"unSubscriptionListId,omitempty,string"`
	UnsubscriptionLandingPageId       int    `json:"unsubscriptionLandingPageId,omitempty,string"`
}

// Create a new email group in eloqua
// During testing subscriptionLandingPageId & subscriptionLandingPageId seemed to be required but
// as this is not as per the documentation it is not required in this method.
// If you get ObjectValidationError's it may be due to this.
func (e *EmailGroupService) Create(name string, emailGroup *EmailGroup) (*EmailGroup, *Response, error) {
	if emailGroup == nil {
		emailGroup = &EmailGroup{}
	}

	emailGroup.Name = name
	endpoint := "/assets/email/group"
	resp, err := e.client.postRequestDecode(endpoint, emailGroup)
	return emailGroup, resp, err
}

// Get a email group object via its ID
func (e *EmailGroupService) Get(id int) (*EmailGroup, *Response, error) {
	endpoint := fmt.Sprintf("/assets/email/group/%d?depth=complete", id)
	emailGroup := &EmailGroup{}
	resp, err := e.client.getRequestDecode(endpoint, emailGroup)
	return emailGroup, resp, err
}

// List many eloqua email groups
func (e *EmailGroupService) List(opts *ListOptions) ([]EmailGroup, *Response, error) {
	endpoint := "/assets/email/groups"
	emailGroups := new([]EmailGroup)
	resp, err := e.client.getRequestListDecode(endpoint, emailGroups, opts)
	return *emailGroups, resp, err
}

// Update an existing email group in eloqua
// During testing subscriptionLandingPageId & subscriptionLandingPageId seemed to be required but
// as this is not as per the documentation it is not required in this method.
// If you get ObjectValidationError's it may be due to this.
func (e *EmailGroupService) Update(id int, name string, emailGroup *EmailGroup) (*EmailGroup, *Response, error) {
	if emailGroup == nil {
		emailGroup = &EmailGroup{}
	}

	emailGroup.ID = id
	emailGroup.Name = name

	endpoint := fmt.Sprintf("/assets/email/group/%d", emailGroup.ID)
	resp, err := e.client.putRequestDecode(endpoint, emailGroup)
	return emailGroup, resp, err
}

// Delete an existing email group from eloqua
func (e *EmailGroupService) Delete(id int) (*Response, error) {
	emailGroup := &EmailGroup{ID: id}
	endpoint := fmt.Sprintf("/assets/email/group/%d", emailGroup.ID)
	resp, err := e.client.deleteRequest(endpoint, emailGroup)
	return resp, err
}
