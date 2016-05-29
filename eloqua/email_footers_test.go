package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestEmailFooterCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailFooter{Name: "A Test Footer"}

	addRestHandlerFunc("/assets/email/footer", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(EmailFooter)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailFooter.Create body", v, input)

		fmt.Fprint(w, `{"type":"EmailFooter","id":"10005","name":"A Test Footer"}`)
	})

	emailFooter, _, err := client.EmailFooters.Create("A Test Footer", nil)
	if err != nil {
		t.Errorf("EmailFooters.Create recieved error: %v", err)
	}

	output := &EmailFooter{ID: 10005, Name: "A Test Footer", Type: "EmailFooter"}

	testModels(t, "EmailFooters.Create", emailFooter, output)
}

func TestEmailFooterGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/email/footer/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"EmailFooter","id":"10005","name":"A Test Footer", "folderId": "101","images": [{"name":"test image","type":"ImageFile","size":{"type":"Size","width":"100","height":"117"}}]}`)
	})

	emailFooter, _, err := client.EmailFooters.Get(1005)
	if err != nil {
		t.Errorf("EmailFooters.Get recieved error: %v", err)
	}

	output := &EmailFooter{ID: 10005, Name: "A Test Footer", FolderID: 101, Images: []Image{
		Image{Type: "ImageFile", Name: "test image", Size: Size{Type: "Size", Width: 100, Height: 117}},
	}}
	testModels(t, "EmailFooters.Get", emailFooter, output)
}

func TestEmailFooterList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/email/footers", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"EmailFooter","id":"10005","name":"A Test Footer"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	emailFooters, resp, err := client.EmailFooters.List(reqOpts)
	if err != nil {
		t.Errorf("EmailFooters.List recieved error: %v", err)
	}

	want := []EmailFooter{{Type: "EmailFooter", ID: 10005, Name: "A Test Footer"}}
	testModels(t, "EmailFooters.List", emailFooters, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("EmailFooters.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("EmailFooters.List response page number incorrect")
	}
}

func TestEmailFooterUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailFooter{ID: 10005, Name: "Updated Footer", Body: "<test-html>"}

	addRestHandlerFunc("/assets/email/footer/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(EmailFooter)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailFooters.Update body", v, input)

		fmt.Fprintf(w, `{"type":"EmailFooter","id":"10005","name":"%s","body":"<test-html>"}`, v.Name)
	})

	emailFooter, _, err := client.EmailFooters.Update(10005, "Updated Footer", input)
	if err != nil {
		t.Errorf("EmailFooters.Update recieved error: %v", err)
	}

	testModels(t, "EmailFooters.Update", emailFooter, input)
}

func TestEmailFooterUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailFooter{ID: 10005, Name: "Updated Footer"}

	addRestHandlerFunc("/assets/email/footer/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(EmailFooter)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailFooters.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"EmailFooter","id":"10005","name":"%s","body":"<test-html>"}`, v.Name)
	})

	emailFooter, _, err := client.EmailFooters.Update(10005, "Updated Footer", nil)
	if err != nil {
		t.Errorf("EmailFooters.Update(Without Model) recieved error: %v", err)
	}

	input.Body = "<test-html>"
	input.Type = "EmailFooter"

	testModels(t, "EmailFooters.Update(Without Model)", emailFooter, input)
}

func TestEmailFooterDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailFooter{ID: 10005}

	addRestHandlerFunc("/assets/email/footer/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(EmailFooter)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailFooters.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.EmailFooters.Delete(10005)
	if err != nil {
		t.Errorf("EmailFooters.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("EmailFooters.Delete request failed")
	}
}
