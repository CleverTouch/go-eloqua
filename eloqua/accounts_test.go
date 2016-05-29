package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestAccountCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Account{Name: "Test Account", PostalCode: "AB12 3CD"}

	addRestHandlerFunc("/data/account", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Account)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Account.Create body", v, input)

		fmt.Fprint(w, `{"assetType":"Account","id":"5","name":"Test Account", "postalCode": "AB12 3CD"}`)
	})

	account, _, err := client.Accounts.Create("Test Account", input)
	if err != nil {
		t.Errorf("Accounts.Create recieved error: %v", err)
	}

	input.ID = 5

	testModels(t, "Accounts.Create", account, input)
}

func TestAccountCreateWithoutModel(t *testing.T) {
	setup()
	defer teardown()

	input := &Account{Name: "Test Account"}

	addRestHandlerFunc("/data/account", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Account)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Account.Create body", v, input)

		fmt.Fprint(w, `{"assetType":"Account","id":"5","name":"Test Account","emailAddress":"Test Account", "postalCode": "AB12 3CD"}`)
	})

	account, _, err := client.Accounts.Create("Test Account", nil)
	if err != nil {
		t.Errorf("Accounts.Create recieved error: %v", err)
	}

	input.ID = 5
	input.PostalCode = "AB12 3CD"
	input.Name = "Test Account"

	testModels(t, "Accounts.Create", account, input)
}

func TestAccountGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/data/account/1", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		rJSON := `{"type":"Account", "id": "1", "name":"Test Account 1"}`
		fmt.Fprint(w, rJSON)
	})

	account, _, err := client.Accounts.Get(1)
	if err != nil {
		t.Errorf("Accounts.Get recieved error: %v", err)
	}

	want := &Account{ID: 1, Name: "Test Account 1", Type: "Account"}
	testModels(t, "Accounts.Get", account, want)
}

func TestAccountList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/data/accounts", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "200")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"id":"100", "name":"Test account 100","type": "Account"}], "page":1,"pageSize":200,"total":2}`
		fmt.Fprint(w, rJSON)
	})

	accounts, resp, err := client.Accounts.List(reqOpts)
	if err != nil {
		t.Errorf("Accounts.List recieved error: %v", err)
	}

	want := []Account{{ID: 100, Name: "Test account 100", Type: "Account"}}
	testModels(t, "Accounts.List", accounts, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("Accounts.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("Accounts.List response page number incorrect")
	}
}

func TestAccountUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Account{Name: "Test Account", ID: 2, Country: "United Kingdom"}

	addRestHandlerFunc("/data/account/2", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Account)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Accounts.Update body", v, input)

		fmt.Fprintf(w, `{"type":"Account","id":"2","Name":"%s","country":"United Kingdom"}`, v.Name)
	})

	account, _, err := client.Accounts.Update(2, "Test Account", input)
	if err != nil {
		t.Errorf("Accounts.Update recieved error: %v", err)
	}

	input.Name = "Test Account Updated"

	testModels(t, "Accounts.Update", account, input)
}

func TestAccountUpdateWithoutModel(t *testing.T) {
	setup()
	defer teardown()

	input := &Account{Name: "Test Account", ID: 2}

	addRestHandlerFunc("/data/account/2", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Account)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Accounts.Update body", v, input)

		fmt.Fprintf(w, `{"type":"Account","id":"2","name":"%s","country":"United Kingdom"}`, v.Name)
	})

	account, _, err := client.Accounts.Update(2, "Test Account", nil)
	if err != nil {
		t.Errorf("Accounts.Update recieved error: %v", err)
	}

	input.Name = "Test Account"
	input.Country = "United Kingdom"
	input.Type = "Account"

	testModels(t, "Accounts.Update", account, input)
}

func TestAccountDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &Account{ID: 5}

	addRestHandlerFunc("/data/account/5", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(Account)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Accounts.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.Accounts.Delete(5)
	if err != nil {
		t.Errorf("Accounts.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Accounts.Delete request failed")
	}
}
