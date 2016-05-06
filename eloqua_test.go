package eloqua

import (
	"net/http"
	"net/http/httptest"
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

func testUrlParam(t *testing.T, req *http.Request, name string, expectedVal string) {
	recievedVal := req.URL.Query().Get(name)
	if recievedVal != expectedVal {
		t.Errorf("URL parameter '%s' is %s, expected %s", name, recievedVal, expectedVal)
	}
}

func TestAuthHeader(t *testing.T) {
	setup()
	defer teardown()

	expectedString := "Basic VGVzdENvbXBhbnlcSm9obi5TbWl0aDpteXNlY3JldA=="

	if client.authHeader != expectedString {
		t.Errorf("Auth header is not as expected \nExpected: %s \nRecieved: %s", expectedString, client.authHeader)
	}
}
