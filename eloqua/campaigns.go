package eloqua

import (
	"fmt"
)

// CampaignService provides access to all the endpoints related
// to campaign data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/2.0 Endpoints/Campaigns/campaigns-API.htm
type CampaignService struct {
	client *Client
}

// Campaign represents an Eloqua campaign object.
// Campaigns are often in other Eloqua models such as Emails & Landing Pages
type Campaign struct {
	Type          string   `json:"type,omitempty"`
	CurrentStatus string   `json:"currentStatus,omitempty"`
	ID            int      `json:"id,omitempty,string"`
	CreatedAt     int      `json:"createdAt,omitempty,string"`
	CreatedBy     int      `json:"createdBy,omitempty,string"`
	Depth         string   `json:"depth,omitempty"`
	Description   string   `json:"description,omitempty"`
	FolderID      int      `json:"folderId,omitempty,string"`
	Name          string   `json:"name,omitempty"`
	Permissions   []string `json:"permissions,omitempty"`
	UpdatedAt     int      `json:"updatedAt,omitempty,string"`
	UpdatedBy     int      `json:"updatedBy,omitempty,string"`

	Elements     []CampaignElement `json:"elements,omitempty"`
	ActualCost   float32           `json:"actualCost,omitempty,string"`
	BudgetedCost float32           `json:"budgetedCost,omitempty,string"`
	CampaignType string            `json:"campaignType,omitempty"`
	FieldValues  []FieldValue      `json:"fieldValues,omitempty"`

	IsEmailMarketingCampaign bool   `json:"isEmailMarketingCampaign,omitempty,string"`
	IsMemberAllowedReEntry   bool   `json:"isMemberAllowedReEntry,omitempty,string"`
	IsReadOnly               bool   `json:"isReadOnly,omitempty,string"`
	IsIncludedInROI          bool   `json:"isIncludedInROI,omitempty,string"`
	IsSyncedWithCRM          bool   `json:"isSyncedWithCRM,omitempty,string"`
	RunAsUserID              int    `json:"runAsUserId,omitempty,string"`
	EndAt                    int    `json:"endAt,omitempty,string"`
	MemberCount              int    `json:"memberCount,omitempty,string"`
	CRMId                    string `json:"crmId,omitempty"`
	Product                  string `json:"product,omitempty"`
	Region                   string `json:"region,omitempty"`
	CampaignCategory         string `json:"campaignCategory,omitempty"`
}

// CampaignElement represents a generic Eloqua campaign step.
// Steps do have their own action-specific properties but, for simplicity, only the common
// properties are used below.
type CampaignElement struct {
	Type            string                   `json:"type,omitempty"`
	ID              int                      `json:"id,omitempty,string"`
	Name            string                   `json:"name,omitempty"`
	MemberCount     int                      `json:"memberCount,omitempty,string"`
	OutputTerminals []CampaignOutputTerminal `json:"outputTerminals,omitempty"`
	Position        Position                 `json:"position,omitempty"`
}

// CampaignOutputTerminal represents the output flows of an element on a campaign.
type CampaignOutputTerminal struct {
	Type          string `json:"type,omitempty"`
	ID            int    `json:"id,omitempty,string"`
	ConnectedID   int    `json:"connectedId,omitempty,string"`
	ConnectedType string `json:"connectedType,omitempty"`
	TerminalType  string `json:"terminalType,omitempty"`
}

// Create a new campaign in eloqua
func (e *CampaignService) Create(name string, campaign *Campaign) (*Campaign, *Response, error) {
	if campaign == nil {
		campaign = &Campaign{}
	}
	campaign.Name = name

	endpoint := "/assets/campaign"
	resp, err := e.client.postRequestDecode(endpoint, campaign)
	return campaign, resp, err
}

// Get an campaign object via its ID
func (e *CampaignService) Get(id int) (*Campaign, *Response, error) {
	endpoint := fmt.Sprintf("/assets/campaign/%d?depth=complete", id)
	campaign := &Campaign{}
	resp, err := e.client.getRequestDecode(endpoint, campaign)
	return campaign, resp, err
}

// List many eloqua campaigns
func (e *CampaignService) List(opts *ListOptions) ([]Campaign, *Response, error) {
	endpoint := "/assets/campaigns"
	campaigns := new([]Campaign)
	resp, err := e.client.getRequestListDecode(endpoint, campaigns, opts)
	return *campaigns, resp, err
}

// Update an existing campaign in eloqua
func (e *CampaignService) Update(id int, name string, campaign *Campaign) (*Campaign, *Response, error) {
	if campaign == nil {
		campaign = &Campaign{}
	}

	campaign.ID = id
	campaign.Name = name

	endpoint := fmt.Sprintf("/assets/campaign/%d", campaign.ID)
	resp, err := e.client.putRequestDecode(endpoint, campaign)
	return campaign, resp, err
}

// Delete an existing campaign from eloqua
func (e *CampaignService) Delete(id int) (*Response, error) {
	campaign := &Campaign{ID: id}
	endpoint := fmt.Sprintf("/assets/campaign/%d", campaign.ID)
	resp, err := e.client.deleteRequest(endpoint, campaign)
	return resp, err
}
