package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestEmailGroupCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailGroup{Name: "My new email group"}

	addRestHandlerFunc("/assets/email/group", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(EmailGroup)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailGroup.Create body", v, input)

		fmt.Fprint(w, `{"type":"EmailGroup","id":"55","name":"My new email group","CreatedAt": "1463510360"}`)
	})

	emailGroup, _, err := client.EmailGroups.Create("My new email group", nil)
	if err != nil {
		t.Errorf("EmailGroups.Create recieved error: %v", err)
	}

	output := &EmailGroup{Type: "EmailGroup", ID: 55, Name: "My new email group", CreatedAt: 1463510360}
	testModels(t, "EmailGroups.Create", emailGroup, output)
}

func TestEmailGroupGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/email/group/55", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"type":"EmailGroup","id":"55","name":"My new email group","createdAt": "1463510360", "description": "Event list"}`)
	})

	emailGroup, _, err := client.EmailGroups.Get(55)
	if err != nil {
		t.Errorf("EmailGroups.Get recieved error: %v", err)
	}

	output := &EmailGroup{Type: "EmailGroup", ID: 55, Name: "My new email group", CreatedAt: 1463510360, Description: "Event list"}
	testModels(t, "EmailGroups.Get", emailGroup, output)
}

func TestEmailGroupListing(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/assets/email/groups", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "200")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"EmailGroup","id":"55","name":"My new email group","createdAt": "1463510360", "description": "Event list"}], "page":1,"pageSize":200,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	emailGroups, resp, err := client.EmailGroups.List(reqOpts)
	if err != nil {
		t.Errorf("EmailGroups.List recieved error: %v", err)
	}

	want := []EmailGroup{{Type: "EmailGroup", ID: 55, Name: "My new email group", CreatedAt: 1463510360, Description: "Event list"}}
	testModels(t, "EmailGroups.List", emailGroups, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("EmailGroups.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("EmailGroups.List response page number incorrect")
	}
}

func TestEmailGroupUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailGroup{ID: 55, Description: "My New Group"}

	addRestHandlerFunc("/assets/email/group/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(EmailGroup)
		json.NewDecoder(req.Body).Decode(v)
		input.Name = "Custom List 1"
		testModels(t, "EmailGroups.Update body", v, input)

		fmt.Fprintf(w, `{"type":"EmailGroup","id":"55","name":"%s","description": "My New Group"}`, v.Name)
	})

	emailGroup, _, err := client.EmailGroups.Update(55, "Custom List 1", input)
	if err != nil {
		t.Errorf("EmailGroups.Update recieved error: %v", err)
	}

	input.Type = "EmailGroup"

	testModels(t, "EmailGroups.Update", emailGroup, input)
}

func TestEmailGroupUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailGroup{ID: 55, Name: "Custom List 1"}

	addRestHandlerFunc("/assets/email/group/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(EmailGroup)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailGroups.Update body", v, input)

		fmt.Fprintf(w, `{"type":"EmailGroup","id":"55","name":"%s","description": "My New Group"}`, v.Name)
	})

	emailGroup, _, err := client.EmailGroups.Update(55, "Custom List 1", nil)
	if err != nil {
		t.Errorf("EmailGroups.Update recieved error: %v", err)
	}
	input.Description = "My New Group"
	input.Type = "EmailGroup"

	testModels(t, "EmailGroups.Update", emailGroup, input)
}

func TestEmailGroupDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailGroup{ID: 55}

	addRestHandlerFunc("/assets/email/group/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(EmailGroup)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailGroups.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.EmailGroups.Delete(55)
	if err != nil {
		t.Errorf("EmailGroups.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("EmailGroups.Delete request failed")
	}
}
