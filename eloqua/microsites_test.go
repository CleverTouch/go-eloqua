package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestMicrositeCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Microsite{Name: "A Test Microsite"}

	addRestHandlerFunc("/assets/microsite", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Microsite)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Microsite.Create body", v, input)

		fmt.Fprint(w, `{"type":"Microsite","id":"10005","name":"A Test Microsite"}`)
	})

	microsite, _, err := client.Microsites.Create("A Test Microsite", nil)
	if err != nil {
		t.Errorf("Microsites.Create recieved error: %v", err)
	}

	output := &Microsite{ID: 10005, Name: "A Test Microsite", Type: "Microsite"}

	testModels(t, "Microsites.Create", microsite, output)
}

func TestMicrositeGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/microsite/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"Microsite","id":"10005","name":"A Test Microsite", "updatedAt": "1329842061","domains":["test.com", "example.com"]}`)
	})

	microsite, _, err := client.Microsites.Get(1005)
	if err != nil {
		t.Errorf("Microsites.Get recieved error: %v", err)
	}

	output := &Microsite{ID: 10005, Name: "A Test Microsite", UpdatedAt: 1329842061, Domains: []string{"test.com", "example.com"}}
	testModels(t, "Microsites.Get", microsite, output)
}

func TestMicrositeList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/microsites", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"Microsite","id":"10005","name":"A Test Microsite"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	microsites, resp, err := client.Microsites.List(reqOpts)
	if err != nil {
		t.Errorf("Microsites.List recieved error: %v", err)
	}

	want := []Microsite{{Type: "Microsite", ID: 10005, Name: "A Test Microsite"}}
	testModels(t, "Microsites.List", microsites, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("Microsites.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("Microsites.List response page number incorrect")
	}
}

func TestMicrositeUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Microsite{ID: 10005, Name: "Updated Microsite", UpdatedAt: 1329842061}

	addRestHandlerFunc("/assets/microsite/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Microsite)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Microsites.Update body", v, input)

		fmt.Fprintf(w, `{"type":"Microsite","id":"10005","name":"%s","updatedAt":"1329842061"}`, v.Name)
	})

	microsite, _, err := client.Microsites.Update(10005, "Updated Microsite", input)
	if err != nil {
		t.Errorf("Microsites.Update recieved error: %v", err)
	}

	testModels(t, "Microsites.Update", microsite, input)
}

func TestMicrositeUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &Microsite{ID: 10005, Name: "Updated Microsite"}

	addRestHandlerFunc("/assets/microsite/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Microsite)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Microsites.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"Microsite","id":"10005","name":"%s","isAuthenticated":"true", "enableWebTrackingOptIn": "disabled"}`, v.Name)
	})

	microsite, _, err := client.Microsites.Update(10005, "Updated Microsite", nil)
	if err != nil {
		t.Errorf("Microsites.Update(Without Model) recieved error: %v", err)
	}

	input.EnableWebTrackingOptIn = "disabled"
	input.IsAuthenticated = true
	input.Type = "Microsite"

	testModels(t, "Microsites.Update(Without Model)", microsite, input)
}

func TestMicrositeDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &Microsite{ID: 10005}

	addRestHandlerFunc("/assets/microsite/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(Microsite)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Microsites.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.Microsites.Delete(10005)
	if err != nil {
		t.Errorf("Microsites.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Microsites.Delete request failed")
	}
}
