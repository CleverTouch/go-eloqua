package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestContactFieldCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactField{Name: "First Name", DataType: "text", DisplayType: "text", UpdateType: "newNotBlank"}

	addRestHandlerFunc("/assets/contact/field", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(ContactField)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactField.Create body", v, input)

		fmt.Fprint(w, `{"assetType":"ContactField","id":"10005","name":"First Name","dataType": "text", "displayType": "text", "updateType": "newNotBlank"}`)
	})

	contactField, _, err := client.ContactFields.Create("First Name", "text", "text", "newNotBlank", nil)
	if err != nil {
		t.Errorf("ContactFields.Create recieved error: %v", err)
	}

	output := &ContactField{ID: 10005, Name: "First Name", DataType: "text", DisplayType: "text", UpdateType: "newNotBlank"}

	testModels(t, "ContactFields.Create", contactField, output)
}

func TestContactFieldGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/contact/field/1005", func(w http.ResponseWriter, req *http.Request) {
		testUrlParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"ContactField","id":"10005","name":"First Name","dataType": "text", "displayType": "text", "updateType": "newNotBlank"}`)
	})

	contactField, _, err := client.ContactFields.Get(1005)
	if err != nil {
		t.Errorf("ContactFields.Get recieved error: %v", err)
	}

	output := &ContactField{ID: 10005, Name: "First Name", DataType: "text", DisplayType: "text", UpdateType: "newNotBlank"}
	testModels(t, "ContactFields.Get", contactField, output)
}

func TestContactFieldList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/assets/contact/fields", func(w http.ResponseWriter, req *http.Request) {
		testUrlParam(t, req, "depth", "minimal")
		testUrlParam(t, req, "count", "200")
		testUrlParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJson := `{"elements":[{"assetType":"ContactField","id":"10005","name":"First Name","dataType": "text", "displayType": "text", "updateType": "newNotBlank"}], "page":1,"pageSize":200,"total":1}`
		fmt.Fprint(w, rJson)
	})

	contactFields, resp, err := client.ContactFields.List(reqOpts)
	if err != nil {
		t.Errorf("ContactFields.List recieved error: %v", err)
	}

	want := []ContactField{{ID: 10005, Name: "First Name", DataType: "text", DisplayType: "text", UpdateType: "newNotBlank"}}
	testModels(t, "ContactFields.List", contactFields, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("ContactFields.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("ContactFields.List response page number incorrect")
	}
}

func TestContactFieldUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactField{ID: 10005, Name: "Last Name", DataType: "text", DisplayType: "text", UpdateType: "newNotBlank", IsRequired: true}

	addRestHandlerFunc("/assets/contact/field/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ContactField)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactFields.Update body", v, input)

		fmt.Fprintf(w, `{"assetType":"ContactField","id":"10005","name":"%s","dataType":"text","displayType":"text","updateType":"newNotBlank","isRequired":"true"}`, v.Name)
	})

	contactField, _, err := client.ContactFields.Update(10005, "Last Name", "text", "text", "newNotBlank", &ContactField{IsRequired: true})
	if err != nil {
		t.Errorf("ContactFields.Update recieved error: %v", err)
	}

	input.IsRequired = true

	testModels(t, "ContactFields.Update", contactField, input)
}

func TestContactFieldDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactField{ID: 10005}

	addRestHandlerFunc("/assets/contact/field/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(ContactField)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactFields.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.ContactFields.Delete(10005)
	if err != nil {
		t.Errorf("ContactFields.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("ContactFields.Delete request failed")
	}
}
