package eloqua

import (
	"fmt"
)

// ExternalActivityService provides access to all the endpoints related
// to External Activity data within eloqua.
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/2.0 Endpoints/External activities/get-data-externalActivity.htm
type ExternalActivityService struct {
	client *Client
}

// ExternalActivity represents an Eloqua External Activity object.
type ExternalActivity struct {
	Type         string `json:"type,omitempty"`
	ID           int    `json:"id,omitempty,string"`
	Depth        string `json:"depth,omitempty"`
	Name         string `json:"name,omitempty"`
	ActivityDate int    `json:"activityDate,omitempty,string"`
	ActivityType string `json:"activityType,omitempty"`
	AssetName    string `json:"assetName,omitempty"`
	AssetType    string `json:"assetType,omitempty"`
	ContactID    int    `json:"contactId,omitempty"`
	CampaignID   int    `json:"campaignId,omitempty"`
}

// Create a new External Activity in eloqua.
// This method call is long due to the amount of required fields at this endpoint.
// Although ActivityDate is not required for creation, you should pass it through on the final parameter
// as eloqua will not set this automatically as the current time.
func (e *ExternalActivityService) Create(name string, assetName string, assetType string, activityType string,
	campaignID int, contactID int, externalActivity *ExternalActivity) (*ExternalActivity, *Response, error) {
	if externalActivity == nil {
		externalActivity = &ExternalActivity{}
	}
	externalActivity.Name = name
	externalActivity.AssetName = assetName
	externalActivity.AssetType = assetType
	externalActivity.ActivityType = activityType
	externalActivity.CampaignID = campaignID
	externalActivity.ContactID = contactID

	endpoint := "/data/activity"
	resp, err := e.client.postRequestDecode(endpoint, externalActivity)
	return externalActivity, resp, err
}

// Get an externalActivity object via its ID
func (e *ExternalActivityService) Get(id int) (*ExternalActivity, *Response, error) {
	endpoint := fmt.Sprintf("/data/activity/%d?depth=complete", id)
	externalActivity := &ExternalActivity{}
	resp, err := e.client.getRequestDecode(endpoint, externalActivity)
	return externalActivity, resp, err
}
