package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestFormCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Form{Name: "A Test Form"}

	addRestHandlerFunc("/assets/form", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Form)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Form.Create body", v, input)

		fmt.Fprint(w, `{"type":"Form","id":"10005","name":"A Test Form"}`)
	})

	form, _, err := client.Forms.Create("A Test Form", nil)
	if err != nil {
		t.Errorf("Forms.Create recieved error: %v", err)
	}

	output := &Form{ID: 10005, Name: "A Test Form", Type: "Form"}

	testModels(t, "Forms.Create", form, output)
}

func TestFormGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/form/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"type":"Form","id":"10005","name":"A Test Form", "folderId": "101", "elements": [{"type":"FormField","id":"1030","name":"First Name","instructions":"","style":"{\"fieldSize\":\"large\",\"labelPosition\":\"top\"}","createdFromContactFieldId":"100002","dataType":"text","displayType":"text","fieldMergeId":"50","htmlName":"firstName","validations":[{"type":"FieldValidation","id":"2546","depth":"complete","description":"Form Field Validation Rule","name":"Form Field Validation Rule","condition":{"type":"IsRequiredCondition"},"isEnabled":"true","message":"This field is required"}]}]}`)
	})

	form, _, err := client.Forms.Get(1005)
	if err != nil {
		t.Errorf("Forms.Get recieved error: %v", err)
	}

	output := &Form{Type: "Form", ID: 10005, Name: "A Test Form", FolderID: 101, FormFields: []FormField{
		FormField{
			Type:         "FormField",
			ID:           1030,
			Name:         "First Name",
			Instructions: "",
			Style:        "{\"fieldSize\":\"large\",\"labelPosition\":\"top\"}",
			CreatedFromContactFieldID: 100002,

			DataType:     "text",
			DisplayType:  "text",
			FieldMergeID: 50,
			HTMLName:     "firstName",
			Validations: []FieldValidation{
				FieldValidation{
					Type:        "FieldValidation",
					ID:          2546,
					Depth:       "complete",
					Description: "Form Field Validation Rule",
					Name:        "Form Field Validation Rule",
					Condition: TypeObject{
						Type: "IsRequiredCondition",
					},
					IsEnabled: true,
					Message:   "This field is required",
				},
			},
		},
	}}
	testModels(t, "Forms.Get", form, output)
}

func TestFormList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/forms", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"Form","id":"10005","name":"A Test Form"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	forms, resp, err := client.Forms.List(reqOpts)
	if err != nil {
		t.Errorf("Forms.List recieved error: %v", err)
	}

	want := []Form{{Type: "Form", ID: 10005, Name: "A Test Form"}}
	testModels(t, "Forms.List", forms, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("Forms.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("Forms.List response page number incorrect")
	}
}

func TestFormUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Form{ID: 10005, Name: "Updated Form"}

	addRestHandlerFunc("/assets/form/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Form)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Forms.Update body", v, input)

		fmt.Fprintf(w, `{"type":"Form","id":"10005","name":"%s"}`, v.Name)
	})

	form, _, err := client.Forms.Update(10005, "Updated Form", input)
	if err != nil {
		t.Errorf("Forms.Update recieved error: %v", err)
	}

	testModels(t, "Forms.Update", form, input)
}

func TestFormUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &Form{ID: 10005, Name: "Updated Form"}

	addRestHandlerFunc("/assets/form/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Form)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Forms.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"Form","id":"10005","name":"%s"}`, v.Name)
	})

	form, _, err := client.Forms.Update(10005, "Updated Form", nil)
	if err != nil {
		t.Errorf("Forms.Update(Without Model) recieved error: %v", err)
	}

	input.Type = "Form"

	testModels(t, "Forms.Update(Without Model)", form, input)
}

func TestFormDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &Form{ID: 10005}

	addRestHandlerFunc("/assets/form/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(Form)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Forms.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.Forms.Delete(10005)
	if err != nil {
		t.Errorf("Forms.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Forms.Delete request failed")
	}
}
