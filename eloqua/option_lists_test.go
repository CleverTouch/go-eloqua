package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestOptionListCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &OptionList{Name: "A Test OptionList"}

	addRestHandlerFunc("/assets/optionList", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(OptionList)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "OptionList.Create body", v, input)

		fmt.Fprint(w, `{"type":"OptionList","id":"10005","name":"A Test OptionList"}`)
	})

	optionList, _, err := client.OptionLists.Create("A Test OptionList", nil)
	if err != nil {
		t.Errorf("OptionLists.Create recieved error: %v", err)
	}

	output := &OptionList{ID: 10005, Name: "A Test OptionList", Type: "OptionList"}

	testModels(t, "OptionLists.Create", optionList, output)
}

func TestOptionListGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/optionList/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"OptionList","id":"10005","name":"A Test OptionList","permissions":["Retrieve", "SetSecurity"]}`)
	})

	optionList, _, err := client.OptionLists.Get(1005)
	if err != nil {
		t.Errorf("OptionLists.Get recieved error: %v", err)
	}

	output := &OptionList{ID: 10005, Name: "A Test OptionList", Permissions: []string{"Retrieve", "SetSecurity"}}
	testModels(t, "OptionLists.Get", optionList, output)
}

func TestOptionListList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/optionLists", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"OptionList","id":"10005","name":"A Test OptionList"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	optionLists, resp, err := client.OptionLists.List(reqOpts)
	if err != nil {
		t.Errorf("OptionLists.List recieved error: %v", err)
	}

	want := []OptionList{{Type: "OptionList", ID: 10005, Name: "A Test OptionList"}}
	testModels(t, "OptionLists.List", optionLists, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("OptionLists.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("OptionLists.List response page number incorrect")
	}
}

func TestOptionListUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &OptionList{ID: 10005, Name: "Updated OptionList", Elements: []Option{Option{Type: "Option", DisplayName: "Option 1", Value: "1"}}}

	addRestHandlerFunc("/assets/optionList/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(OptionList)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "OptionLists.Update body", v, input)

		fmt.Fprintf(w, `{"type":"OptionList","id":"10005","name":"%s","elements":[{"type":"Option","displayName":"Option 1","value":"1"}]}`, v.Name)
	})

	optionList, _, err := client.OptionLists.Update(10005, "Updated OptionList", input)
	if err != nil {
		t.Errorf("OptionLists.Update recieved error: %v", err)
	}

	testModels(t, "OptionLists.Update", optionList, input)
}

func TestOptionListUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &OptionList{ID: 10005, Name: "Updated OptionList"}

	addRestHandlerFunc("/assets/optionList/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(OptionList)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "OptionLists.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"OptionList","id":"10005","name":"%s","depth":"complete"}`, v.Name)
	})

	optionList, _, err := client.OptionLists.Update(10005, "Updated OptionList", nil)
	if err != nil {
		t.Errorf("OptionLists.Update(Without Model) recieved error: %v", err)
	}

	input.Depth = "complete"
	input.Type = "OptionList"

	testModels(t, "OptionLists.Update(Without Model)", optionList, input)
}

func TestOptionListDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &OptionList{ID: 10005}

	addRestHandlerFunc("/assets/optionList/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(OptionList)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "OptionLists.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.OptionLists.Delete(10005)
	if err != nil {
		t.Errorf("OptionLists.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("OptionLists.Delete request failed")
	}
}
