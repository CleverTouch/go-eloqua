package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestContactSegmentCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactSegment{Name: "A Test Segment"}

	addRestHandlerFunc("/assets/contact/segment", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(ContactSegment)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactSegment.Create body", v, input)

		fmt.Fprint(w, `{"type":"ContactSegment","id":"10005","name":"A Test Segment"}`)
	})

	contactSegment, _, err := client.ContactSegments.Create("A Test Segment", nil)
	if err != nil {
		t.Errorf("ContactSegments.Create recieved error: %v", err)
	}

	output := &ContactSegment{ID: 10005, Name: "A Test Segment", Type: "ContactSegment"}

	testModels(t, "ContactSegments.Create", contactSegment, output)
}

func TestContactSegmentGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/contact/segment/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"ContactSegment","id":"10005","name":"A Test Segment", "count": "80084"}`)
	})

	contactSegment, _, err := client.ContactSegments.Get(1005)
	if err != nil {
		t.Errorf("ContactSegments.Get recieved error: %v", err)
	}

	output := &ContactSegment{ID: 10005, Name: "A Test Segment", Count: 80084}
	testModels(t, "ContactSegments.Get", contactSegment, output)
}

func TestContactSegmentList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/contact/segments", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"ContactSegment","id":"10005","name":"A Test Segment"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	contactSegments, resp, err := client.ContactSegments.List(reqOpts)
	if err != nil {
		t.Errorf("ContactSegments.List recieved error: %v", err)
	}

	want := []ContactSegment{{Type: "ContactSegment", ID: 10005, Name: "A Test Segment"}}
	testModels(t, "ContactSegments.List", contactSegments, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("ContactSegments.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("ContactSegments.List response page number incorrect")
	}
}

func TestContactSegmentUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactSegment{ID: 10005, Name: "Updated Segment", Description: "A test description"}

	addRestHandlerFunc("/assets/contact/segment/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ContactSegment)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactSegments.Update body", v, input)

		fmt.Fprintf(w, `{"type":"ContactSegment","id":"10005","name":"%s","description":"A test description"}`, v.Name)
	})

	contactSegment, _, err := client.ContactSegments.Update(10005, "Updated Segment", input)
	if err != nil {
		t.Errorf("ContactSegments.Update recieved error: %v", err)
	}

	testModels(t, "ContactSegments.Update", contactSegment, input)
}

func TestContactSegmentUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactSegment{ID: 10005, Name: "Updated Segment"}

	addRestHandlerFunc("/assets/contact/segment/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ContactSegment)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactSegments.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"ContactSegment","id":"10005","name":"%s","description":"A test description"}`, v.Name)
	})

	contactSegment, _, err := client.ContactSegments.Update(10005, "Updated Segment", nil)
	if err != nil {
		t.Errorf("ContactSegments.Update(Without Model) recieved error: %v", err)
	}

	input.Description = "A test description"
	input.Type = "ContactSegment"

	testModels(t, "ContactSegments.Update(Without Model)", contactSegment, input)
}

func TestContactSegmentDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactSegment{ID: 10005}

	addRestHandlerFunc("/assets/contact/segment/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(ContactSegment)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactSegments.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.ContactSegments.Delete(10005)
	if err != nil {
		t.Errorf("ContactSegments.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("ContactSegments.Delete request failed")
	}
}
