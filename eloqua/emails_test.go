package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestEmailCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Email{Name: "Test Email 2", Subject: "A test email"}

	addRestHandlerFunc("/assets/email", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Email)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Email.Create body", v, input)

		fmt.Fprint(w, `{"assetType":"Email","id":"2","Name":"Test Email 2","subject":"A test email"}`)
	})

	email, _, err := client.Emails.Create("Test Email 2", input)
	if err != nil {
		t.Errorf("Emails.Create recieved error: %v", err)
	}

	input.ID = 2

	testModels(t, "Emails.Create", email, input)
}

func TestEmailCreateWithoutPassingEmail(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/email", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Email)
		json.NewDecoder(req.Body).Decode(v)
		expected := &Email{Name: "Test Email 2"}
		testModels(t, "Email.Create body (without model)", v, expected)

		fmt.Fprint(w, `{"assetType":"Email","id":"2","Name":"Test Email 2","subject":"A test email"}`)
	})

	email, _, err := client.Emails.Create("Test Email 2", nil)
	if err != nil {
		t.Errorf("Emails.Create recieved error: %v", err)
	}

	expectedResult := &Email{AssetType: "Email", ID: 2, Name: "Test Email 2", Subject: "A test email"}
	testModels(t, "Emails.Create (without model)", email, expectedResult)
}

func TestEmailGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/email/1", func(w http.ResponseWriter, req *http.Request) {
		testUrlParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		rJson := `{"assetType":"Email", "id": "1", "name":"Test Email 1"}`
		fmt.Fprint(w, rJson)
	})

	email, _, err := client.Emails.Get(1)
	if err != nil {
		t.Errorf("Emails.Get recieved error: %v", err)
	}

	want := &Email{ID: 1, Name: "Test Email 1", AssetType: "Email"}
	testModels(t, "Emails.Get", email, want)
}

func TestEmailList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/assets/emails", func(w http.ResponseWriter, req *http.Request) {
		testUrlParam(t, req, "depth", "minimal")
		testUrlParam(t, req, "count", "200")
		testUrlParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJson := `{"elements":[{"id":"100", "name":"Test email 100","assetType": "Email"}], "page":1,"pageSize":200,"total":2}`
		fmt.Fprint(w, rJson)
	})

	emails, resp, err := client.Emails.List(reqOpts)
	if err != nil {
		t.Errorf("Emails.List recieved error: %v", err)
	}

	want := []Email{{ID: 100, Name: "Test email 100", AssetType: "Email"}}
	testModels(t, "Emails.List", emails, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("Emails.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("Emails.List response page number incorrect")
	}
}

func TestEmailUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Email{Name: "Test Email 2", ID: 2, Subject: "A test email"}

	addRestHandlerFunc("/assets/email/2", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Email)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Emails.Update body", v, input)

		fmt.Fprintf(w, `{"assetType":"Email","id":"2","Name":"%s","subject":"A test email"}`, v.Name)
	})

	email, _, err := client.Emails.Update(2, "Test Email Updated", input)
	if err != nil {
		t.Errorf("Emails.Update recieved error: %v", err)
	}

	input.Name = "Test Email Updated"

	testModels(t, "Emails.Update", email, input)
}

func TestUserUpdateWithoutPassingEmail(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/email/8", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Email)
		json.NewDecoder(req.Body).Decode(v)
		expectedData := &Email{ID: 8, Name: "Test Email Updated"}
		testModels(t, "Emails.Update body (without model)", v, expectedData)
		fmt.Fprintf(w, `{"assetType":"Email","id":"8","Name":"%s","htmlContent":{"type": "RawHtmlContent","html":"Hello"}}`, v.Name)
	})

	email, _, err := client.Emails.Update(8, "Test Email Updated", nil)
	if err != nil {
		t.Errorf("Emails.Update recieved error: %v", err)
	}

	resultModel := &Email{AssetType: "Email", Name: "Test Email Updated",
		ID: 8, HtmlContent: htmlContent{ContentType: "RawHtmlContent", Html: "Hello"}}
	testModels(t, "Emails.Update (without model)", email, resultModel)
}

func TestEmailDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &Email{ID: 2}

	addRestHandlerFunc("/assets/email/2", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(Email)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Emails.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.Emails.Delete(2)
	if err != nil {
		t.Errorf("Emails.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Emails.Delete request failed")
	}
}
