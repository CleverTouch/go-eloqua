package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCustomObjectCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &CustomObject{Name: "My new custom object"}

	addRestHandlerFunc("/assets/customObject", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(CustomObject)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "CustomObject.Create body", v, input)

		fmt.Fprint(w, `{"type":"CustomObject","id":"55","name":"My new custom object","CreatedAt": "1463510360"}`)
	})

	customObject, _, err := client.CustomObjects.Create("My new custom object", nil)
	if err != nil {
		t.Errorf("CustomObjects.Create recieved error: %v", err)
	}

	output := &CustomObject{Type: "CustomObject", ID: 55, Name: "My new custom object", CreatedAt: 1463510360}
	testModels(t, "CustomObjects.Create", customObject, output)
}

func TestCustomObjectGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/customObject/55", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"type":"CustomObject","id":"55","name":"My new custom object","createdAt": "1463510360", "description": "Event list"}`)
	})

	customObject, _, err := client.CustomObjects.Get(55)
	if err != nil {
		t.Errorf("CustomObjects.Get recieved error: %v", err)
	}

	output := &CustomObject{Type: "CustomObject", ID: 55, Name: "My new custom object", CreatedAt: 1463510360, Description: "Event list"}
	testModels(t, "CustomObjects.Get", customObject, output)
}

func TestCustomObjectListing(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/assets/customObjects", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "200")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"CustomObject","id":"55","name":"My new custom object","createdAt": "1463510360", "description": "Event list"}], "page":1,"pageSize":200,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	customObjects, resp, err := client.CustomObjects.List(reqOpts)
	if err != nil {
		t.Errorf("CustomObjects.List recieved error: %v", err)
	}

	want := []CustomObject{{Type: "CustomObject", ID: 55, Name: "My new custom object", CreatedAt: 1463510360, Description: "Event list"}}
	testModels(t, "CustomObjects.List", customObjects, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("CustomObjects.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("CustomObjects.List response page number incorrect")
	}
}

func TestCustomObjectUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &CustomObject{
		ID:          55,
		Description: "My New Object",
		Fields: []CustomObjectField{
			CustomObjectField{Type: "CustomObjectField", ID: 141, Name: "Custom Object Field Example"},
		},
	}

	addRestHandlerFunc("/assets/customObject/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(CustomObject)
		json.NewDecoder(req.Body).Decode(v)
		input.Name = "Custom Object 1"
		testModels(t, "CustomObjects.Update body", v, input)

		fmt.Fprintf(w, `{"type":"CustomObject","id":"55","name":"%s","description": "My New Object",
			"fields": [{"type": "CustomObjectField", "id": "141", "name":"Custom Object Field Example"}]}`, v.Name)
	})

	customObject, _, err := client.CustomObjects.Update(55, "Custom Object 1", input)
	if err != nil {
		t.Errorf("CustomObjects.Update recieved error: %v", err)
	}

	input.Type = "CustomObject"

	testModels(t, "CustomObjects.Update", customObject, input)
}

func TestCustomObjectUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &CustomObject{ID: 55, Name: "Custom Object 1"}

	addRestHandlerFunc("/assets/customObject/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(CustomObject)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "CustomObjects.Update body", v, input)

		fmt.Fprintf(w, `{"type":"CustomObject","id":"55","name":"%s","description": "My New Object"}`, v.Name)
	})

	customObject, _, err := client.CustomObjects.Update(55, "Custom Object 1", nil)
	if err != nil {
		t.Errorf("CustomObjects.Update recieved error: %v", err)
	}
	input.Description = "My New Object"
	input.Type = "CustomObject"

	testModels(t, "CustomObjects.Update", customObject, input)
}

func TestCustomObjectDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &CustomObject{ID: 55}

	addRestHandlerFunc("/assets/customObject/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(CustomObject)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "CustomObjects.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.CustomObjects.Delete(55)
	if err != nil {
		t.Errorf("CustomObjects.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("CustomObjects.Delete request failed")
	}
}
