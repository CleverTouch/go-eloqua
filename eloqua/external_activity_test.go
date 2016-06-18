package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestExternalActivityCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalActivity{Name: "Test External Activity", ActivityType: "visit", AssetName: "/animals/cats", CampaignID: 5, ContactID: 20, AssetType: "website"}

	addRestHandlerFunc("/data/activity", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(ExternalActivity)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalActivity.Create body", v, input)

		fmt.Fprint(w, `{"type":"ExternalActivities","id":"354","depth":"complete","name":"Test External Activity","activityDate":"1461038400","activityType":"visit","assetName":"/animals/cats","assetType":"website","campaignId":5,"contactId":20}`)
	})

	externalActivity, _, err := client.ExternalActivity.Create("Test External Activity", "/animals/cats", "website", "visit", 5, 20, nil)
	if err != nil {
		t.Errorf("ExternalActivity.Create recieved error: %v", err)
	}

	output := &ExternalActivity{Type: "ExternalActivities", ID: 354, Depth: "complete", Name: "Test External Activity", ActivityType: "visit", ActivityDate: 1461038400, AssetName: "/animals/cats", CampaignID: 5, ContactID: 20, AssetType: "website"}

	testModels(t, "ExternalActivity.Create", externalActivity, output)
}

func TestExternalActivityGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/data/activity/354", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"type":"ExternalActivities","id":"354","depth":"complete","name":"Test External Activity","activityDate":"1461038400","activityType":"visit","assetName":"/animals/cats","assetType":"website","campaignId":5,"contactId":20}`)
	})

	externalActivity, _, err := client.ExternalActivity.Get(354)
	if err != nil {
		t.Errorf("ExternalActivity.Get recieved error: %v", err)
	}

	output := &ExternalActivity{Type: "ExternalActivities", ID: 354, Depth: "complete", Name: "Test External Activity", ActivityType: "visit", ActivityDate: 1461038400, AssetName: "/animals/cats", CampaignID: 5, ContactID: 20, AssetType: "website"}
	testModels(t, "ExternalActivity.Get", externalActivity, output)
}
