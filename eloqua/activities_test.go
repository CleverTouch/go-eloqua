package eloqua

import (
	"fmt"
	"net/http"
	"testing"
)

func TestActivityGet(t *testing.T) {
	setup()
	defer teardown()

	addLegacyRestHandlerFunc("/data/activities/contact/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "type", "formSubmit")
		testURLParam(t, req, "startDate", "0")
		testURLParam(t, req, "endDate", "1467051883")
		testURLParam(t, req, "count", "500")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `[{"type":"Activity","activityDate":"1466525000","activityType":"formSubmit","asset":"5","assetType":"form","contact":"1005","details":[{"Key":"Collection","Value":"test mapping"},{"Key":"FormName","Value":"test_form_name"}],"id":"10"}]`)
	})

	activities, _, err := client.Activities.List(1005, "formSubmit", 0, 1467051883, 500)
	if err != nil {
		t.Errorf("Activities.Get recieved error: %v", err)
	}

	if len(activities) != 1 {
		t.Error("Activities received is not a list")
	}

	output := Activity{ID: 10, Type: "Activity", ActivityDate: 1466525000, ActivityType: "formSubmit", Asset: 5, AssetType: "form", Contact: 1005, Details: []ActivityDetail{
		ActivityDetail{Key: "Collection", Value: "test mapping"},
		ActivityDetail{Key: "FormName", Value: "test_form_name"},
	}}
	testModels(t, "Activities.Get", activities[0], output)
}
