package eloqua

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Eloqua client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup creates a test server and client instance to test against.
func setup() {
	// Create test servers
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	// Create Eloqua test client
	client = NewClient(server.URL, "TestCompany", "John.Smith", "mysecret")
}

func addRestHandlerFunc(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	mux.HandleFunc("/api/rest/2.0/"+strings.Trim(endpoint, " /"), handler)
}

// teardown does any cleanup operations, Specifically closes down the http server
func teardown() {
	server.Close()
}

// testURLParam is a helper to check url parameters are as expected
func testURLParam(t *testing.T, req *http.Request, name string, expectedVal string) {
	receivedVal := req.URL.Query().Get(name)
	if receivedVal != expectedVal {
		t.Errorf("URL parameter '%s' is %s, expected %s", name, receivedVal, expectedVal)
	}
}

// testMethod ensures the http method for a request is as expected
func testMethod(t *testing.T, req *http.Request, method string) {
	if req.Method != strings.ToUpper(method) {
		t.Errorf("HTTP method is not as expected\nExpected: %s\nReceived: %s", strings.ToUpper(method), req.Method)
	}
}

// testModels tests two structs against eachother
func testModels(t *testing.T, testDesc string, test interface{}, expected interface{}) {
	if !reflect.DeepEqual(test, expected) {
		t.Errorf("%s data is not as expected.\nReturned \n%+v,\nWanted \n%+v", testDesc, test, expected)
	}
}

func TestAuthHeader(t *testing.T) {
	setup()
	defer teardown()

	expectedString := "Basic VGVzdENvbXBhbnlcSm9obi5TbWl0aDpteXNlY3JldA=="

	if client.authHeader != expectedString {
		t.Errorf("Auth header is not as expected \nExpected: %s \nReceived: %s", expectedString, client.authHeader)
	}
}

func TestRequestDecodeJSONErrorHandling(t *testing.T) {
	setup()
	defer teardown()

	tMap := make(chan int)
	_, err := client.requestDecode("/test/endpoint", "POST", tMap)

	if err.Error() != "json: unsupported type: chan int" {
		t.Error("Delete request with invalid postdata not returning an error as expected")
	}
}

func TestDeleteRequestErrorHandling(t *testing.T) {
	setup()
	defer teardown()

	user := User{Name: "Test User"}
	_, err := client.deleteRequest("/test/endpoint", user)

	if err == nil {
		t.Error("Request did not return an error but a 404 was expected")
	}
}

func TestDeleteRequestJSONErrorHandling(t *testing.T) {
	setup()
	defer teardown()

	tMap := make(chan int)
	_, err := client.deleteRequest("/test/endpoint", tMap)

	if err.Error() != "json: unsupported type: chan int" {
		t.Error("Delete request with invalid postdata not returning an error as expected")
	}
}

func TestEloquaErrorResponse(t *testing.T) {
	setup()
	defer teardown()

	responseMessage := "This is a test error message string response"

	addRestHandlerFunc("/assets/contact/lists", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(401)
		fmt.Fprint(w, responseMessage)
	})

	_, resp, err := client.ContactLists.List(nil)

	if resp.StatusCode != 401 {
		t.Errorf("Wrong response status code, Expected %d, Received %d", 401, resp.StatusCode)
	}

	if err.Error() != errorMessages[401] {
		t.Errorf("Wrong error message received, \nExpected: %s\nRecieved: %s", errorMessages[401], err.Error())
	}

	if resp.ErrorContent != responseMessage {
		t.Errorf("Failed request content not in request body as expected.\nExpected: %s\nRecieved:%s", responseMessage, resp.ErrorContent)
	}
}

func TestEloquaDefaultErrorResponse(t *testing.T) {
	setup()
	defer teardown()

	responseMessage := "There was an issue performing your request"

	addRestHandlerFunc("/assets/contact/lists", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(608)
		fmt.Fprint(w, responseMessage)
	})

	_, resp, _ := client.ContactLists.List(nil)

	if resp.ErrorContent != responseMessage {
		t.Errorf("Failed request content not in request body as expected.\nExpected: %s\nRecieved:%s", responseMessage, resp.ErrorContent)
	}
}
