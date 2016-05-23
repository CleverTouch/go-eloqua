package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestContentSectionCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &ContentSection{Name: "My new content section"}

	addRestHandlerFunc("/assets/contentSection", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(ContentSection)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContentSection.Create body", v, input)

		fmt.Fprint(w, `{"type":"ContentSection","id":"55","name":"My new content section","CreatedAt": "1463510360"}`)
	})

	contentSection, _, err := client.ContentSections.Create("My new content section", nil)
	if err != nil {
		t.Errorf("ContentSections.Create recieved error: %v", err)
	}

	output := &ContentSection{Type: "ContentSection", ID: 55, Name: "My new content section", CreatedAt: 1463510360}
	testModels(t, "ContentSections.Create", contentSection, output)
}

func TestContentSectionGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/contentSection/55", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"type":"ContentSection","id":"55","name":"My new content section","createdAt": "1463510360", "description": "Event list"}`)
	})

	contentSection, _, err := client.ContentSections.Get(55)
	if err != nil {
		t.Errorf("ContentSections.Get recieved error: %v", err)
	}

	output := &ContentSection{Type: "ContentSection", ID: 55, Name: "My new content section", CreatedAt: 1463510360, Description: "Event list"}
	testModels(t, "ContentSections.Get", contentSection, output)
}

func TestContentSectionListing(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/assets/contentSections", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "200")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"ContentSection","id":"55","name":"My new content section","createdAt": "1463510360", "description": "Event list"}], "page":1,"pageSize":200,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	contentSections, resp, err := client.ContentSections.List(reqOpts)
	if err != nil {
		t.Errorf("ContentSections.List recieved error: %v", err)
	}

	want := []ContentSection{{Type: "ContentSection", ID: 55, Name: "My new content section", CreatedAt: 1463510360, Description: "Event list"}}
	testModels(t, "ContentSections.List", contentSections, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("ContentSections.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("ContentSections.List response page number incorrect")
	}
}

func TestContentSectionUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &ContentSection{ID: 55, Description: "My New List"}

	addRestHandlerFunc("/assets/contentSection/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ContentSection)
		json.NewDecoder(req.Body).Decode(v)
		input.Name = "Custom List 1"
		testModels(t, "ContentSections.Update body", v, input)

		fmt.Fprintf(w, `{"type":"ContentSection","id":"55","name":"%s","description": "My New List"}`, v.Name)
	})

	contentSection, _, err := client.ContentSections.Update(55, "Custom List 1", input)
	if err != nil {
		t.Errorf("ContentSections.Update recieved error: %v", err)
	}

	input.Type = "ContentSection"

	testModels(t, "ContentSections.Update", contentSection, input)
}

func TestContentSectionUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &ContentSection{ID: 55, Name: "Custom List 1"}

	addRestHandlerFunc("/assets/contentSection/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ContentSection)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContentSections.Update body", v, input)

		fmt.Fprintf(w, `{"type":"ContentSection","id":"55","name":"%s","description": "My New List"}`, v.Name)
	})

	contentSection, _, err := client.ContentSections.Update(55, "Custom List 1", nil)
	if err != nil {
		t.Errorf("ContentSections.Update recieved error: %v", err)
	}
	input.Description = "My New List"
	input.Type = "ContentSection"

	testModels(t, "ContentSections.Update", contentSection, input)
}

func TestContentSectionDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &ContentSection{ID: 55}

	addRestHandlerFunc("/assets/contentSection/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(ContentSection)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ContentSections.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.ContentSections.Delete(55)
	if err != nil {
		t.Errorf("ContentSections.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("ContentSections.Delete request failed")
	}
}
