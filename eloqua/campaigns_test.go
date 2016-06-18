package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCampaignCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Campaign{Name: "A Test Campaign"}

	addRestHandlerFunc("/assets/campaign", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Campaign)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Campaign.Create body", v, input)

		fmt.Fprint(w, `{"type":"Campaign","id":"10005","name":"A Test Campaign"}`)
	})

	campaign, _, err := client.Campaigns.Create("A Test Campaign", nil)
	if err != nil {
		t.Errorf("Campaigns.Create recieved error: %v", err)
	}

	output := &Campaign{ID: 10005, Name: "A Test Campaign", Type: "Campaign"}

	testModels(t, "Campaigns.Create", campaign, output)
}

func TestCampaignGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/campaign/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"Campaign","id":"10005","name":"A Test Campaign","updatedAt":"1329842061","elements":[{"type":"CampaignSegment","id":"4440","name":"Segment Members","memberCount":"0","outputTerminals":[{"type":"CampaignOutputTerminal","id":"4430","connectedId":"4441","connectedType":"CampaignEmail","terminalType":"out"}],"position":{"type":"Position","x":"382","y":"136"},"isFinished":"true","isRecurring":"false","segmentId":"352"}]}`)
	})

	campaign, _, err := client.Campaigns.Get(1005)
	if err != nil {
		t.Errorf("Campaigns.Get recieved error: %v", err)
	}

	output := &Campaign{ID: 10005, Name: "A Test Campaign", UpdatedAt: 1329842061, Elements: []CampaignElement{
		CampaignElement{
			Type:        "CampaignSegment",
			ID:          4440,
			Name:        "Segment Members",
			MemberCount: 0,
			OutputTerminals: []CampaignOutputTerminal{
				CampaignOutputTerminal{Type: "CampaignOutputTerminal", ID: 4430, ConnectedID: 4441, ConnectedType: "CampaignEmail", TerminalType: "out"},
			},
			Position: Position{Type: "Position", X: 382, Y: 136},
		},
	}}
	testModels(t, "Campaigns.Get", campaign, output)
}

func TestCampaignList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/campaigns", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"Campaign","id":"10005","name":"A Test Campaign"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	campaigns, resp, err := client.Campaigns.List(reqOpts)
	if err != nil {
		t.Errorf("Campaigns.List recieved error: %v", err)
	}

	want := []Campaign{{Type: "Campaign", ID: 10005, Name: "A Test Campaign"}}
	testModels(t, "Campaigns.List", campaigns, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("Campaigns.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("Campaigns.List response page number incorrect")
	}
}

func TestCampaignUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Campaign{ID: 10005, Name: "Updated Campaign", UpdatedAt: 1329842061}

	addRestHandlerFunc("/assets/campaign/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Campaign)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Campaigns.Update body", v, input)

		fmt.Fprintf(w, `{"type":"Campaign","id":"10005","name":"%s","updatedAt":"1329842061","fieldValues":[{"type":"FieldValue","id":"7","value":"Test Field Value"},{"type":"FieldValue","id":"8","value":"Email Tracking Test Value"}]}`, v.Name)
	})

	campaign, _, err := client.Campaigns.Update(10005, "Updated Campaign", input)
	if err != nil {
		t.Errorf("Campaigns.Update recieved error: %v", err)
	}

	input.FieldValues = []FieldValue{
		FieldValue{Type: "FieldValue", ID: 7, Value: "Test Field Value"},
		FieldValue{Type: "FieldValue", ID: 8, Value: "Email Tracking Test Value"},
	}

	testModels(t, "Campaigns.Update", campaign, input)
}

func TestCampaignUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &Campaign{ID: 10005, Name: "Updated Campaign"}

	addRestHandlerFunc("/assets/campaign/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Campaign)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Campaigns.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"Campaign","id":"10005","name":"%s","isSyncedWithCRM": "true","memberCount":"300"}`, v.Name)
	})

	campaign, _, err := client.Campaigns.Update(10005, "Updated Campaign", nil)
	if err != nil {
		t.Errorf("Campaigns.Update(Without Model) recieved error: %v", err)
	}

	input.Type = "Campaign"
	input.IsSyncedWithCRM = true
	input.MemberCount = 300

	testModels(t, "Campaigns.Update(Without Model)", campaign, input)
}

func TestCampaignDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &Campaign{ID: 10005}

	addRestHandlerFunc("/assets/campaign/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(Campaign)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Campaigns.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.Campaigns.Delete(10005)
	if err != nil {
		t.Errorf("Campaigns.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Campaigns.Delete request failed")
	}
}
