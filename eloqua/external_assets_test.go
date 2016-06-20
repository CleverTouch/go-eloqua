package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestExternalAssetCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalAsset{Name: "A Test ExternalAsset"}

	addRestHandlerFunc("/assets/external", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(ExternalAsset)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalAsset.Create body", v, input)

		fmt.Fprint(w, `{"type":"ExternalAsset","id":"10005","name":"A Test ExternalAsset"}`)
	})

	externalAsset, _, err := client.ExternalAssets.Create("A Test ExternalAsset", nil)
	if err != nil {
		t.Errorf("ExternalAssets.Create recieved error: %v", err)
	}

	output := &ExternalAsset{ID: 10005, Name: "A Test ExternalAsset", Type: "ExternalAsset"}

	testModels(t, "ExternalAssets.Create", externalAsset, output)
}

func TestExternalAssetGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/external/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"ExternalAsset","id":"10005","name":"A Test ExternalAsset","updatedAt":"1329842061","externalAssetTypeId": "10"}`)
	})

	externalAsset, _, err := client.ExternalAssets.Get(1005)
	if err != nil {
		t.Errorf("ExternalAssets.Get recieved error: %v", err)
	}

	output := &ExternalAsset{ID: 10005, Name: "A Test ExternalAsset", UpdatedAt: 1329842061, ExternalAssetTypeID: 10}
	testModels(t, "ExternalAssets.Get", externalAsset, output)
}

func TestExternalAssetList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/externals", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"ExternalAsset","id":"10005","name":"A Test ExternalAsset"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	externalAssets, resp, err := client.ExternalAssets.List(reqOpts)
	if err != nil {
		t.Errorf("ExternalAssets.List recieved error: %v", err)
	}

	want := []ExternalAsset{{Type: "ExternalAsset", ID: 10005, Name: "A Test ExternalAsset"}}
	testModels(t, "ExternalAssets.List", externalAssets, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("ExternalAssets.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("ExternalAssets.List response page number incorrect")
	}
}

func TestExternalAssetUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalAsset{ID: 10005, Name: "Updated ExternalAsset", UpdatedAt: 1329842061}

	addRestHandlerFunc("/assets/external/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ExternalAsset)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalAssets.Update body", v, input)

		fmt.Fprintf(w, `{"type":"ExternalAsset","id":"10005","name":"%s","updatedAt":"1329842061"}`, v.Name)
	})

	externalAsset, _, err := client.ExternalAssets.Update(10005, "Updated ExternalAsset", input)
	if err != nil {
		t.Errorf("ExternalAssets.Update recieved error: %v", err)
	}

	testModels(t, "ExternalAssets.Update", externalAsset, input)
}

func TestExternalAssetUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalAsset{ID: 10005, Name: "Updated ExternalAsset"}

	addRestHandlerFunc("/assets/external/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(ExternalAsset)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalAssets.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"ExternalAsset","id":"10005","name":"%s"}`, v.Name)
	})

	externalAsset, _, err := client.ExternalAssets.Update(10005, "Updated ExternalAsset", nil)
	if err != nil {
		t.Errorf("ExternalAssets.Update(Without Model) recieved error: %v", err)
	}

	input.Type = "ExternalAsset"

	testModels(t, "ExternalAssets.Update(Without Model)", externalAsset, input)
}

func TestExternalAssetDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &ExternalAsset{ID: 10005}

	addRestHandlerFunc("/assets/external/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(ExternalAsset)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "ExternalAssets.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.ExternalAssets.Delete(10005)
	if err != nil {
		t.Errorf("ExternalAssets.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("ExternalAssets.Delete request failed")
	}
}
