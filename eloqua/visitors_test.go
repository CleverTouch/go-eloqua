package eloqua

import (
	"fmt"
	"net/http"
	"testing"
)

func TestVisitorList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 50, Page: 1}

	addRestHandlerFunc("/data/visitors", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "50")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"Visitor","id":"10005","contactId": "6","visitorId": "10", "V_LastVisitDateAndTime": "1464545297"}], "page":1,"pageSize":50,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	visitors, resp, err := client.Visitors.List(reqOpts)
	if err != nil {
		t.Errorf("Visitors.List recieved error: %v", err)
	}

	want := []Visitor{{Type: "Visitor", VisitorID: 10, ContactID: 6, LastVisitDateAndTime: 1464545297}}
	testModels(t, "Visitors.List", visitors, want)

	if resp.PageSize != reqOpts.Count {
		t.Errorf("Visitors.List response page size incorrect.\nExpected: %d\nRecieved: %d", reqOpts.Count, resp.PageSize)
	}
	if resp.Page != reqOpts.Page {
		t.Error("Visitors.List response page number incorrect")
	}
}
