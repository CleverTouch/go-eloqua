package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestExternalAssetTypeCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalAssetType{Name: "A Test ExternalAssetType"}

	addRestHandlerFunc("/assets/external/type", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(ExternalAssetType)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalAssetType.Create body", v, input)

		fmt.Fprint(w, `{"type":"ExternalAssetType","id":"10005","name":"A Test ExternalAssetType"}`)
	})

	externalAssetType, _, err := client.ExternalAssetTypes.Create("A Test ExternalAssetType", nil)
	if err != nil {
		t.Errorf("ExternalAssetTypes.Create recieved error: %v", err)
	}

	output := &ExternalAssetType{ID: 10005, Name: "A Test ExternalAssetType", Type: "ExternalAssetType"}

	testModels(t, "ExternalAssetTypes.Create", externalAssetType, output)
}

func TestExternalAssetTypeGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/external/type/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"ExternalAssetType","id":"10005","name":"A Test ExternalAssetType","updatedAt":"1329842061","activityTypes":[{"type":"ExternalActivityType","id":"1007","createdAt":"1466250500","createdBy":"56","depth":"complete","name":"form-submit","updatedAt":"1466250500","updatedBy":"56"}]}`)
	})

	externalAssetType, _, err := client.ExternalAssetTypes.Get(1005)
	if err != nil {
		t.Errorf("ExternalAssetTypes.Get recieved error: %v", err)
	}

	output := &ExternalAssetType{ID: 10005, Name: "A Test ExternalAssetType", UpdatedAt: 1329842061, ActivityTypes: []ExternalActivityType{
		ExternalActivityType{Type: "ExternalActivityType", ID: 1007, CreatedAt: 1466250500, UpdatedAt: 1466250500, Depth: "complete", Name: "form-submit", CreatedBy: 56, UpdatedBy: 56},
	}}
	testModels(t, "ExternalAssetTypes.Get", externalAssetType, output)
}

func TestExternalAssetTypeList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/external/types", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"ExternalAssetType","id":"10005","name":"A Test ExternalAssetType"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	externalAssetTypes, resp, err := client.ExternalAssetTypes.List(reqOpts)
	if err != nil {
		t.Errorf("ExternalAssetTypes.List recieved error: %v", err)
	}

	want := []ExternalAssetType{{Type: "ExternalAssetType", ID: 10005, Name: "A Test ExternalAssetType"}}
	testModels(t, "ExternalAssetTypes.List", externalAssetTypes, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("ExternalAssetTypes.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("ExternalAssetTypes.List response page number incorrect")
	}
}

func TestExternalAssetTypeUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalAssetType{ID: 10005, Name: "Updated ExternalAssetType", UpdatedAt: 1329842061}

	addRestHandlerFunc("/assets/external/type/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ExternalAssetType)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalAssetTypes.Update body", v, input)

		fmt.Fprintf(w, `{"type":"ExternalAssetType","id":"10005","name":"%s","updatedAt":"1329842061"}`, v.Name)
	})

	externalAssetType, _, err := client.ExternalAssetTypes.Update(10005, "Updated ExternalAssetType", input)
	if err != nil {
		t.Errorf("ExternalAssetTypes.Update recieved error: %v", err)
	}

	testModels(t, "ExternalAssetTypes.Update", externalAssetType, input)
}

func TestExternalAssetTypeUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalAssetType{ID: 10005, Name: "Updated ExternalAssetType"}

	addRestHandlerFunc("/assets/external/type/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ExternalAssetType)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalAssetTypes.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"ExternalAssetType","id":"10005","name":"%s"}`, v.Name)
	})

	externalAssetType, _, err := client.ExternalAssetTypes.Update(10005, "Updated ExternalAssetType", nil)
	if err != nil {
		t.Errorf("ExternalAssetTypes.Update(Without Model) recieved error: %v", err)
	}

	input.Type = "ExternalAssetType"

	testModels(t, "ExternalAssetTypes.Update(Without Model)", externalAssetType, input)
}

func TestExternalAssetTypeDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalAssetType{ID: 10005}

	addRestHandlerFunc("/assets/external/type/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(ExternalAssetType)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalAssetTypes.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.ExternalAssetTypes.Delete(10005)
	if err != nil {
		t.Errorf("ExternalAssetTypes.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("ExternalAssetTypes.Delete request failed")
	}
}
