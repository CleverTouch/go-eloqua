package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestLandingPageCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &LandingPage{Name: "A Test LandingPage"}

	addRestHandlerFunc("/assets/landingPage", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(LandingPage)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "LandingPage.Create body", v, input)

		fmt.Fprint(w, `{"type":"LandingPage","id":"10005","name":"A Test LandingPage"}`)
	})

	landingPage, _, err := client.LandingPages.Create("A Test LandingPage", nil)
	if err != nil {
		t.Errorf("LandingPages.Create recieved error: %v", err)
	}

	output := &LandingPage{ID: 10005, Name: "A Test LandingPage", Type: "LandingPage"}

	testModels(t, "LandingPages.Create", landingPage, output)
}

func TestLandingPageGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/landingPage/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"LandingPage","id":"10005","name":"A Test LandingPage","updatedAt":"1329842061","htmlContent":{"type": "RawHtmlContent", "contentSource": "upload","html":"<p>hello</p>"}}`)
	})

	landingPage, _, err := client.LandingPages.Get(1005)
	if err != nil {
		t.Errorf("LandingPages.Get recieved error: %v", err)
	}

	output := &LandingPage{ID: 10005, Name: "A Test LandingPage", UpdatedAt: 1329842061, HTMLContent: HTMLContent{
		Type:          "RawHtmlContent",
		ContentSource: "upload",
		HTML:          "<p>hello</p>",
	}}
	testModels(t, "LandingPages.Get", landingPage, output)
}

func TestLandingPageList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/landingPages", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"LandingPage","id":"10005","name":"A Test LandingPage"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	landingPages, resp, err := client.LandingPages.List(reqOpts)
	if err != nil {
		t.Errorf("LandingPages.List recieved error: %v", err)
	}

	want := []LandingPage{{Type: "LandingPage", ID: 10005, Name: "A Test LandingPage"}}
	testModels(t, "LandingPages.List", landingPages, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("LandingPages.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("LandingPages.List response page number incorrect")
	}
}

func TestLandingPageUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &LandingPage{ID: 10005, Name: "Updated LandingPage", UpdatedAt: 1329842061}

	addRestHandlerFunc("/assets/landingPage/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(LandingPage)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "LandingPages.Update body", v, input)

		fmt.Fprintf(w, `{"type":"LandingPage","id":"10005","name":"%s","updatedAt":"1329842061"}`, v.Name)
	})

	landingPage, _, err := client.LandingPages.Update(10005, "Updated LandingPage", input)
	if err != nil {
		t.Errorf("LandingPages.Update recieved error: %v", err)
	}

	testModels(t, "LandingPages.Update", landingPage, input)
}

func TestLandingPageUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &LandingPage{ID: 10005, Name: "Updated LandingPage"}

	addRestHandlerFunc("/assets/landingPage/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(LandingPage)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "LandingPages.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"LandingPage","id":"10005","name":"%s", "isContentProtected": "false"}`, v.Name)
	})

	landingPage, _, err := client.LandingPages.Update(10005, "Updated LandingPage", nil)
	if err != nil {
		t.Errorf("LandingPages.Update(Without Model) recieved error: %v", err)
	}

	input.Type = "LandingPage"
	input.IsContentProtected = false

	testModels(t, "LandingPages.Update(Without Model)", landingPage, input)
}

func TestLandingPageDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &LandingPage{ID: 10005}

	addRestHandlerFunc("/assets/landingPage/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(LandingPage)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "LandingPages.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.LandingPages.Delete(10005)
	if err != nil {
		t.Errorf("LandingPages.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("LandingPages.Delete request failed")
	}
}
