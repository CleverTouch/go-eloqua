package eloqua

import (
	"fmt"
)

// ActivityService provides access to all the endpoints related
// to activity data within eloqua.
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Activities/activities-API.htm
type ActivityService struct {
	client *Client
}

// Activity represents an Eloqua activity objects.
type Activity struct {
	Type         string           `json:"type,omitempty"`
	ActivityDate int              `json:"activityDate,omitempty,string"`
	ActivityType string           `json:"activityType,omitempty"`
	Asset        int              `json:"asset,omitempty,string"`
	AssetType    string           `json:"assetType,omitempty"`
	Contact      int              `json:"contact,omitempty,string"`
	ID           int              `json:"id,omitempty,string"`
	Details      []ActivityDetail `json:"details,omitempty"`
}

// ActivityDetail is a detail item that is provided with an activity.
// It is a key-value pair of information relating to the specific activity instance.
type ActivityDetail struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

// List many Eloqua activities.
// Due to this being an old 1.0 endpoint this does not give the usual listing result,
// It will only provide a simple list of activity items.
func (e *ActivityService) List(contactID int, activtyType string, startDate int, endDate int, count int) ([]Activity, *Response, error) {
	queryString := fmt.Sprintf("type=%s&startDate=%d&endDate=%d&count=%d", activtyType, startDate, endDate, count)
	endpoint := fmt.Sprintf("/api/rest/1.0/data/activities/contact/%d?%s", contactID, queryString)
	activities := new([]Activity)
	resp, err := e.client.getRequestDecode(endpoint, activities)
	return *activities, resp, err
}
