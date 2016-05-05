package eloqua

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ = fmt.Printf
var _ = ioutil.ReadAll

type Client struct {
	client *http.Client

	// The base URL for the eloqua instance
	BaseURL string

	// Eloqua login company name
	companyName string
	// Eloqua login user name
	userName string
	// Eloqua login password
	password string
	// Basic auth header value
	authHeader string

	// Various components of the API
	Emails *EmailService
}

// NewClient creates a new instance of an Eloqua HTTP client
// used to interface with the Eloqua API.
func NewClient(baseURL string, companyName string, userName string, password string) *Client {

	authString := companyName + "\\" + userName + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))

	c := &Client{
		client:      http.DefaultClient,
		BaseURL:     strings.Trim(baseURL, " /"),
		companyName: companyName,
		userName:    userName,
		password:    password,
		authHeader:  "Basic " + encodedAuth,
	}

	// Create services
	c.Emails = &EmailService{client: c}

	return c
}

type Response struct {
	*http.Response
}

// newReponse creates a new custom Response for the given http.Response
func newResponse(r *http.Response) *Response {
	return &Response{Response: r}
}

// Perform a request to the Eloqua API
// Flexible to allow any use of any endpoint but it only returns a
// simple respose.
func (c *Client) RestRequest(endpoint string, method string, jsonData string) (*Response, error) {
	url := c.BaseURL + "/api/rest/2.0/" + strings.Trim(endpoint, " /")
	// fmt.Println(jsonData)
	jsonStr := []byte(jsonData)
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.authHeader)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

// Performs a GET request and decodes the response into the provided interface
func (c *Client) getRequestDecode(endpoint string, v interface{}) (*Response, error) {
	resp, err := c.RestRequest(endpoint, "GET", "")
	defer resp.Body.Close()
	// TODO - check the response further for common eloqua responses
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

// Performs a POST request and decodes the response into the provided interface
func (c *Client) postRequestDecode(endpoint string, v interface{}) (*Response, error) {

	postBody := ""

	if v != nil {
		jsonString, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		postBody = string(jsonString)
	}

	resp, err := c.RestRequest(endpoint, "POST", postBody)
	defer resp.Body.Close()
	// TODO - check the response further for common eloqua responses
	if err != nil {
		return resp, err
	}

	// content, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(content))

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
	}

	return resp, err
}
