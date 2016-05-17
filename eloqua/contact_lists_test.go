package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestContactListCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactList{Name: "My new contact list"}

	addRestHandlerFunc("/assets/contact/list", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(ContactList)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactList.Create body", v, input)

		fmt.Fprint(w, `{"type":"ContactList","id":"55","name":"My new contact list","CreatedAt": "1463510360"}`)
	})

	contactList, _, err := client.ContactLists.Create("My new contact list", nil)
	if err != nil {
		t.Errorf("ContactLists.Create recieved error: %v", err)
	}

	output := &ContactList{Type: "ContactList", ID: 55, Name: "My new contact list", CreatedAt: 1463510360}
	testModels(t, "ContactLists.Create", contactList, output)
}

func TestContactListGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/contact/list/55", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"type":"ContactList","id":"55","name":"My new contact list","createdAt": "1463510360", "description": "Event list"}`)
	})

	contactList, _, err := client.ContactLists.Get(55)
	if err != nil {
		t.Errorf("ContactLists.Get recieved error: %v", err)
	}

	output := &ContactList{Type: "ContactList", ID: 55, Name: "My new contact list", CreatedAt: 1463510360, Description: "Event list"}
	testModels(t, "ContactLists.Get", contactList, output)
}

func TestContactListListing(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/assets/contact/lists", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "200")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"ContactList","id":"55","name":"My new contact list","createdAt": "1463510360", "description": "Event list"}], "page":1,"pageSize":200,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	contactLists, resp, err := client.ContactLists.List(reqOpts)
	if err != nil {
		t.Errorf("ContactLists.List recieved error: %v", err)
	}

	want := []ContactList{{Type: "ContactList", ID: 55, Name: "My new contact list", CreatedAt: 1463510360, Description: "Event list"}}
	testModels(t, "ContactLists.List", contactLists, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("ContactLists.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("ContactLists.List response page number incorrect")
	}
}

func TestContactListUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactList{ID: 55, Description: "My New List"}

	addRestHandlerFunc("/assets/contact/list/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ContactList)
		json.NewDecoder(req.Body).Decode(v)
		input.Name = "Custom List 1"
		testModels(t, "ContactLists.Update body", v, input)

		fmt.Fprintf(w, `{"type":"ContactList","id":"55","name":"%s","description": "My New List"}`, v.Name)
	})

	contactList, _, err := client.ContactLists.Update(55, "Custom List 1", input)
	if err != nil {
		t.Errorf("ContactLists.Update recieved error: %v", err)
	}

	input.Type = "ContactList"

	testModels(t, "ContactLists.Update", contactList, input)
}

func TestContactListUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactList{ID: 55, Name: "Custom List 1"}

	addRestHandlerFunc("/assets/contact/list/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ContactList)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactLists.Update body", v, input)

		fmt.Fprintf(w, `{"type":"ContactList","id":"55","name":"%s","description": "My New List"}`, v.Name)
	})

	contactList, _, err := client.ContactLists.Update(55, "Custom List 1", nil)
	if err != nil {
		t.Errorf("ContactLists.Update recieved error: %v", err)
	}
	input.Description = "My New List"
	input.Type = "ContactList"

	testModels(t, "ContactLists.Update", contactList, input)
}

func TestContactListDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &ContactList{ID: 55}

	addRestHandlerFunc("/assets/contact/list/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(ContactList)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContactLists.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.ContactLists.Delete(55)
	if err != nil {
		t.Errorf("ContactLists.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("ContactLists.Delete request failed")
	}
}
