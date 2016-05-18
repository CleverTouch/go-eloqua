package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestEmailFolderCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailFolder{Name: "A Test Folder"}

	addRestHandlerFunc("/assets/email/folder", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(EmailFolder)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailFolder.Create body", v, input)

		fmt.Fprint(w, `{"type":"EmailFolder","id":"10005","name":"A Test Folder"}`)
	})

	emailFolder, _, err := client.EmailFolders.Create("A Test Folder", nil)
	if err != nil {
		t.Errorf("EmailFolders.Create recieved error: %v", err)
	}

	output := &EmailFolder{ID: 10005, Name: "A Test Folder", Type: "EmailFolder"}

	testModels(t, "EmailFolders.Create", emailFolder, output)
}

func TestEmailFolderGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/email/folder/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"EmailFolder","id":"10005","name":"A Test Folder", "folderId": "101"}`)
	})

	emailFolder, _, err := client.EmailFolders.Get(1005)
	if err != nil {
		t.Errorf("EmailFolders.Get recieved error: %v", err)
	}

	output := &EmailFolder{ID: 10005, Name: "A Test Folder", FolderId: 101}
	testModels(t, "EmailFolders.Get", emailFolder, output)
}

func TestEmailFolderList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/email/folders", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"EmailFolder","id":"10005","name":"A Test Folder"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	emailFolders, resp, err := client.EmailFolders.List(reqOpts)
	if err != nil {
		t.Errorf("EmailFolders.List recieved error: %v", err)
	}

	want := []EmailFolder{{Type: "EmailFolder", ID: 10005, Name: "A Test Folder"}}
	testModels(t, "EmailFolders.List", emailFolders, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("EmailFolders.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("EmailFolders.List response page number incorrect")
	}
}

func TestEmailFolderUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailFolder{ID: 10005, Name: "Updated Folder", Description: "A test description"}

	addRestHandlerFunc("/assets/email/folder/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(EmailFolder)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailFolders.Update body", v, input)

		fmt.Fprintf(w, `{"type":"EmailFolder","id":"10005","name":"%s","description":"A test description"}`, v.Name)
	})

	emailFolder, _, err := client.EmailFolders.Update(10005, "Updated Folder", input)
	if err != nil {
		t.Errorf("EmailFolders.Update recieved error: %v", err)
	}

	testModels(t, "EmailFolders.Update", emailFolder, input)
}

func TestEmailFolderUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailFolder{ID: 10005, Name: "Updated Folder"}

	addRestHandlerFunc("/assets/email/folder/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(EmailFolder)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailFolders.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"EmailFolder","id":"10005","name":"%s","description":"A test description"}`, v.Name)
	})

	emailFolder, _, err := client.EmailFolders.Update(10005, "Updated Folder", nil)
	if err != nil {
		t.Errorf("EmailFolders.Update(Without Model) recieved error: %v", err)
	}

	input.Description = "A test description"
	input.Type = "EmailFolder"

	testModels(t, "EmailFolders.Update(Without Model)", emailFolder, input)
}

func TestEmailFolderDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &EmailFolder{ID: 10005}

	addRestHandlerFunc("/assets/email/folder/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(EmailFolder)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "EmailFolders.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.EmailFolders.Delete(10005)
	if err != nil {
		t.Errorf("EmailFolders.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("EmailFolders.Delete request failed")
	}
}
