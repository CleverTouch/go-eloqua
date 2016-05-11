package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestContactCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Contact{Name: "test@example.com", FirstName: "John Smith"}

	addRestHandlerFunc("/data/contact", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Contact)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Contact.Create body", v, input)

		fmt.Fprint(w, `{"assetType":"Contact","id":"5","name":"test@example.com","firstName":"John Smith"}`)
	})

	contact, _, err := client.Contacts.Create("test@example.com", input)
	if err != nil {
		t.Errorf("Contacts.Create recieved error: %v", err)
	}

	input.ID = 5

	testModels(t, "Contacts.Create", contact, input)
}

func TestContactGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/data/contact/1", func(w http.ResponseWriter, req *http.Request) {
		testUrlParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		rJson := `{"type":"Contact", "id": "1", "name":"Test Contact 1"}`
		fmt.Fprint(w, rJson)
	})

	contact, _, err := client.Contacts.Get(1)
	if err != nil {
		t.Errorf("Contacts.Get recieved error: %v", err)
	}

	want := &Contact{ID: 1, Name: "Test Contact 1", Type: "Contact"}
	testModels(t, "Contacts.Get", contact, want)
}

func TestContactList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/data/contacts", func(w http.ResponseWriter, req *http.Request) {
		testUrlParam(t, req, "depth", "minimal")
		testUrlParam(t, req, "count", "200")
		testUrlParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJson := `{"elements":[{"id":"100", "name":"Test contact 100","type": "Contact"}], "page":1,"pageSize":200,"total":2}`
		fmt.Fprint(w, rJson)
	})

	contacts, resp, err := client.Contacts.List(reqOpts)
	if err != nil {
		t.Errorf("Contacts.List recieved error: %v", err)
	}

	want := []Contact{{ID: 100, Name: "Test contact 100", Type: "Contact"}}
	testModels(t, "Contacts.List", contacts, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("Contacts.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("Contacts.List response page number incorrect")
	}
}

func TestContactUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Contact{Name: "Test Contact 2", ID: 2, IsSubscribed: false}

	addRestHandlerFunc("/data/contact/2", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Contact)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Contacts.Update body", v, input)

		fmt.Fprintf(w, `{"type":"Contact","id":"2","Name":"%s","isSubscribed":"false"}`, v.Name)
	})

	contact, _, err := client.Contacts.Update(2, "Test Contact Updated", input)
	if err != nil {
		t.Errorf("Contacts.Update recieved error: %v", err)
	}

	input.Name = "Test Contact Updated"

	testModels(t, "Contacts.Update", contact, input)
}

func TestContactDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &Contact{ID: 5}

	addRestHandlerFunc("/data/contact/5", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(Contact)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Contacts.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.Contacts.Delete(5)
	if err != nil {
		t.Errorf("Contacts.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Contacts.Delete request failed")
	}
}
