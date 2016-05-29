package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestEmailHeaderCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailHeader{Name: "A Test Header"}

	addRestHandlerFunc("/assets/email/header", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(EmailHeader)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailHeader.Create body", v, input)

		fmt.Fprint(w, `{"type":"EmailHeader","id":"10005","name":"A Test Header"}`)
	})

	emailHeader, _, err := client.EmailHeaders.Create("A Test Header", nil)
	if err != nil {
		t.Errorf("EmailHeaders.Create recieved error: %v", err)
	}

	output := &EmailHeader{ID: 10005, Name: "A Test Header", Type: "EmailHeader"}

	testModels(t, "EmailHeaders.Create", emailHeader, output)
}

func TestEmailHeaderGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/email/header/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"EmailHeader","id":"10005","name":"A Test Header", "folderId": "101","images": [{"name":"test image","type":"ImageFile","size":{"type":"Size","width":"100","height":"117"}}]}`)
	})

	emailHeader, _, err := client.EmailHeaders.Get(1005)
	if err != nil {
		t.Errorf("EmailHeaders.Get recieved error: %v", err)
	}

	output := &EmailHeader{ID: 10005, Name: "A Test Header", FolderID: 101, Images: []Image{
		Image{Type: "ImageFile", Name: "test image", Size: Size{Type: "Size", Width: 100, Height: 117}},
	}}
	testModels(t, "EmailHeaders.Get", emailHeader, output)
}

func TestEmailHeaderList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/email/headers", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"EmailHeader","id":"10005","name":"A Test Header"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	emailHeaders, resp, err := client.EmailHeaders.List(reqOpts)
	if err != nil {
		t.Errorf("EmailHeaders.List recieved error: %v", err)
	}

	want := []EmailHeader{{Type: "EmailHeader", ID: 10005, Name: "A Test Header"}}
	testModels(t, "EmailHeaders.List", emailHeaders, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("EmailHeaders.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("EmailHeaders.List response page number incorrect")
	}
}

func TestEmailHeaderUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailHeader{ID: 10005, Name: "Updated Header", Body: "<test-html>"}

	addRestHandlerFunc("/assets/email/header/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(EmailHeader)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailHeaders.Update body", v, input)

		fmt.Fprintf(w, `{"type":"EmailHeader","id":"10005","name":"%s","body":"<test-html>"}`, v.Name)
	})

	emailHeader, _, err := client.EmailHeaders.Update(10005, "Updated Header", input)
	if err != nil {
		t.Errorf("EmailHeaders.Update recieved error: %v", err)
	}

	testModels(t, "EmailHeaders.Update", emailHeader, input)
}

func TestEmailHeaderUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailHeader{ID: 10005, Name: "Updated Header"}

	addRestHandlerFunc("/assets/email/header/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(EmailHeader)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailHeaders.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"EmailHeader","id":"10005","name":"%s","body":"<test-html>"}`, v.Name)
	})

	emailHeader, _, err := client.EmailHeaders.Update(10005, "Updated Header", nil)
	if err != nil {
		t.Errorf("EmailHeaders.Update(Without Model) recieved error: %v", err)
	}

	input.Body = "<test-html>"
	input.Type = "EmailHeader"

	testModels(t, "EmailHeaders.Update(Without Model)", emailHeader, input)
}

func TestEmailHeaderDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailHeader{ID: 10005}

	addRestHandlerFunc("/assets/email/header/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(EmailHeader)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailHeaders.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.EmailHeaders.Delete(10005)
	if err != nil {
		t.Errorf("EmailHeaders.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("EmailHeaders.Delete request failed")
	}
}
