package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestUserGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/system/user/1", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		rJSON := `{"type":"User", "id": "1", "name":"Test User 1"}`
		fmt.Fprint(w, rJSON)
	})

	user, _, err := client.Users.Get(1)
	if err != nil {
		t.Errorf("Users.Get recieved error: %v", err)
	}

	want := &User{ID: 1, Name: "Test User 1", Type: "User"}
	testModels(t, "Users.Get", user, want)
}

func TestUserList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 200, Page: 1}

	addRestHandlerFunc("/system/users", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "200")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"id":"100", "name":"Test user 100","type": "User"}], "page":1,"pageSize":200,"total":2}`
		fmt.Fprint(w, rJSON)
	})

	users, resp, err := client.Users.List(reqOpts)
	if err != nil {
		t.Errorf("Users.List recieved error: %v", err)
	}

	want := []User{{ID: 100, Name: "Test user 100", Type: "User"}}
	testModels(t, "Users.List", users, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("Users.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("Users.List response page number incorrect")
	}
}

func TestUserUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &User{Name: "Test User 2", ID: 2, Description: "A test user"}

	addRestHandlerFunc("/system/user/2", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(User)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Users.Update body", v, input)

		fmt.Fprintf(w, `{"type":"User","id":"2","Name":"%s","description":"A test user"}`, v.Name)
	})

	user, _, err := client.Users.Update(2, "Test User Updated", input)
	if err != nil {
		t.Errorf("Users.Update recieved error: %v", err)
	}

	input.Name = "Test User Updated"
	testModels(t, "Users.Update", user, input)
}

func TestUserUpdateWithoutPassingUser(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/system/user/8", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(User)
		json.NewDecoder(req.Body).Decode(v)
		expectedData := &User{ID: 8, Name: "Test User Updated"}
		testModels(t, "Users.Update body (without passing model)", v, expectedData)
		fmt.Fprintf(w, `{"type":"User","id":"8","Name":"%s","description":"A test user"}`, v.Name)
	})

	user, _, err := client.Users.Update(8, "Test User Updated", nil)
	if err != nil {
		t.Errorf("Users.Update recieved error: %v", err)
	}

	resultModel := &User{Type: "User", Name: "Test User Updated", ID: 8, Description: "A test user"}
	testModels(t, "Users.Update (without passing model)", user, resultModel)
}
