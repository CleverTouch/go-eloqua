package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCustomObjectDataCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &CustomObjectData{}

	addRestHandlerFunc("/data/customObject/55/instance", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(CustomObjectData)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "CustomObjectData.Create body", v, input)

		fmt.Fprint(w, `{"type":"CustomObjectData","id":"3","fieldValues":[{"type":"FieldValue","id":"613","value":"My test value"}]}`)
	})

	customObjectData, _, err := client.CustomObjectData.Create(55, nil)
	if err != nil {
		t.Errorf("CustomObjectData.Create recieved error: %v", err)
	}

	output := &CustomObjectData{Type: "CustomObjectData", ID: 3, FieldValues: []FieldValue{FieldValue{Type: "FieldValue", ID: 613, Value: "My test value"}}}
	testModels(t, "CustomObjectData.Create", customObjectData, output)
}

func TestCustomObjectDataGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/data/customObject/55/instance/10", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"type":"CustomObjectData","id":"3","fieldValues":[{"type":"FieldValue","id":"613","value":"My test value"}]}`)
	})

	customObjectData, _, err := client.CustomObjectData.Get(55, 10)
	if err != nil {
		t.Errorf("CustomObjectData.Get recieved error: %v", err)
	}

	output := &CustomObjectData{Type: "CustomObjectData", ID: 3, FieldValues: []FieldValue{FieldValue{Type: "FieldValue", ID: 613, Value: "My test value"}}}
	testModels(t, "CustomObjectData.Get", customObjectData, output)
}

func TestCustomObjectDataListing(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/data/customObject/55/instances", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "200")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"CustomObjectData","id":"1","createdAt":"1464118483","name":"My test value 3","fieldValues":[{"type":"FieldValue","id":"613","value":"My test value 3"}],"uniqueCode":"DCTET000000000001"}],"page":1,"pageSize":200,"total":3}`
		fmt.Fprint(w, rJSON)
	})

	customObjectDatas, resp, err := client.CustomObjectData.List(55, reqOpts)
	if err != nil {
		t.Errorf("CustomObjectData.List recieved error: %v", err)
	}

	want := []CustomObjectData{CustomObjectData{Type: "CustomObjectData", ID: 1, CreatedAt: 1464118483, Name: "My test value 3", FieldValues: []FieldValue{FieldValue{Type: "FieldValue", ID: 613, Value: "My test value 3"}}, UniqueCode: "DCTET000000000001"}}
	testModels(t, "CustomObjectData.List", customObjectDatas, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("CustomObjectData.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("CustomObjectData.List response page number incorrect")
	}
}

func TestCustomObjectDataUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &CustomObjectData{ID: 10, FieldValues: []FieldValue{
		FieldValue{ID: 613, Value: "My new test value"},
	}}

	addRestHandlerFunc("/data/customObject/55/instance/10", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(CustomObjectData)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "CustomObjectData.Update body", v, input)

		fmt.Fprint(w, `{"type":"CustomObjectData","id":"3","fieldValues":[{"type":"FieldValue","id":"613","value":"My new test value"}]}`)
	})

	customObjectData, _, err := client.CustomObjectData.Update(55, 10, input)
	if err != nil {
		t.Errorf("CustomObjectData.Update recieved error: %v", err)
	}

	input.Type = "CustomObjectData"
	input.FieldValues[0].Type = "FieldValue"

	testModels(t, "CustomObjectData.Update", customObjectData, input)
}

func TestCustomObjectDataUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &CustomObjectData{ID: 10}

	addRestHandlerFunc("/data/customObject/55/instance/10", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(CustomObjectData)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "CustomObjectData.Update body", v, input)

		fmt.Fprint(w, `{"type":"CustomObjectData","id":"10","fieldValues":[{"type":"FieldValue","id":"613","value":"My new test value"}]}`)
	})

	customObjectData, _, err := client.CustomObjectData.Update(55, 10, nil)
	if err != nil {
		t.Errorf("CustomObjectData.Update recieved error: %v", err)
	}

	output := &CustomObjectData{Type: "CustomObjectData", ID: 10, FieldValues: []FieldValue{
		FieldValue{Type: "FieldValue", ID: 613, Value: "My new test value"},
	}}

	testModels(t, "CustomObjectData.Update", customObjectData, output)
}

func TestCustomObjectDataDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &CustomObjectData{ID: 10}

	addRestHandlerFunc("/data/customObject/55/instance/10", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(CustomObjectData)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "CustomObjectData.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.CustomObjectData.Delete(55, 10)
	if err != nil {
		t.Errorf("CustomObjectData.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("CustomObjectData.Delete request failed")
	}
}
