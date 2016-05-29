package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestFormDataCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &FormData{}

	addRestHandlerFunc("/data/form/55", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(FormData)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "FormData.Create body", v, input)

		fmt.Fprint(w, `{"type":"FormData","id":"3","fieldValues":[{"type":"FieldValue","id":"613","value":"My test value"}]}`)
	})

	formData, _, err := client.FormData.Create(55, nil)
	if err != nil {
		t.Errorf("FormData.Create recieved error: %v", err)
	}

	output := &FormData{Type: "FormData", ID: 3, FieldValues: []FieldValue{FieldValue{Type: "FieldValue", ID: 613, Value: "My test value"}}}
	testModels(t, "FormData.Create", formData, output)
}

func TestFormDataListing(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/data/form/55", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "200")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"FormData","id":"1","submittedAt":"1464118483","name":"My test value 3","fieldValues":[{"type":"FieldValue","id":"613","value":"My test value 3"}]}],"page":1,"pageSize":200,"total":3}`
		fmt.Fprint(w, rJSON)
	})

	formDatas, resp, err := client.FormData.List(55, reqOpts)
	if err != nil {
		t.Errorf("FormData.List recieved error: %v", err)
	}

	want := []FormData{FormData{Type: "FormData", ID: 1, SubmittedAt: 1464118483, Name: "My test value 3", FieldValues: []FieldValue{FieldValue{Type: "FieldValue", ID: 613, Value: "My test value 3"}}}}
	testModels(t, "FormData.List", formDatas, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("FormData.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("FormData.List response page number incorrect")
	}
}
