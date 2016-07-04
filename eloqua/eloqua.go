package eloqua

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

var _ = fmt.Printf
var _ = ioutil.ReadAll

// Client manages communications with the Eloqua API. It contains services to access each
// endpoint grouping so the API can be used in a fluent manner.
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

	// The service endpoints of the API
	Accounts           *AccountService
	Activities         *ActivityService
	Campaigns          *CampaignService
	Contacts           *ContactService
	ContactFields      *ContactFieldService
	ContactLists       *ContactListService
	ContactSegments    *ContactSegmentService
	ContentSections    *ContentSectionService
	CustomObjects      *CustomObjectService
	CustomObjectData   *CustomObjectDataService
	Emails             *EmailService
	EmailFolders       *EmailFolderService
	EmailGroups        *EmailGroupService
	EmailHeaders       *EmailHeaderService
	EmailFooters       *EmailFooterService
	ExternalActivity   *ExternalActivityService
	ExternalAssets     *ExternalAssetService
	ExternalAssetTypes *ExternalAssetTypeService
	Forms              *FormService
	FormData           *FormDataService
	Images             *ImageService
	LandingPages       *LandingPageService
	Microsites         *MicrositeService
	OptionLists        *OptionListService
	Users              *UserService
	Visitors           *VisitorService
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
	c.Accounts = &AccountService{client: c}
	c.Activities = &ActivityService{client: c}
	c.Campaigns = &CampaignService{client: c}
	c.Contacts = &ContactService{client: c}
	c.ContactFields = &ContactFieldService{client: c}
	c.ContactLists = &ContactListService{client: c}
	c.ContactSegments = &ContactSegmentService{client: c}
	c.ContentSections = &ContentSectionService{client: c}
	c.CustomObjects = &CustomObjectService{client: c}
	c.CustomObjectData = &CustomObjectDataService{client: c}
	c.Emails = &EmailService{client: c}
	c.EmailFolders = &EmailFolderService{client: c}
	c.EmailGroups = &EmailGroupService{client: c}
	c.EmailHeaders = &EmailHeaderService{client: c}
	c.EmailFooters = &EmailFooterService{client: c}
	c.ExternalActivity = &ExternalActivityService{client: c}
	c.ExternalAssets = &ExternalAssetService{client: c}
	c.ExternalAssetTypes = &ExternalAssetTypeService{client: c}
	c.Forms = &FormService{client: c}
	c.FormData = &FormDataService{client: c}
	c.Images = &ImageService{client: c}
	c.LandingPages = &LandingPageService{client: c}
	c.Microsites = &MicrositeService{client: c}
	c.OptionLists = &OptionListService{client: c}
	c.Users = &UserService{client: c}
	c.Visitors = &VisitorService{client: c}

	return c
}

// Response is a custom http response that, upon a standard http response,
// contains eloqua specific details such as listing properies and error details.
type Response struct {
	*http.Response

	// Variables used in listing operations.
	// Will remain zero-valued for other operations

	// The main body containing the request entities
	Elements json.RawMessage `json:"elements,omitempty"`
	// The current page of the response
	Page int `json:"page,omitempty"`
	// The page size of the response
	PageSize int `json:"pageSize,omitempty"`
	// The total entities found in the query
	Total int `json:"total,omitempty"`

	// The returned response body in the event of an error
	// Use this to help debug in the event of unknown errors
	ErrorContent string `json:"-"`
}

// newReponse creates a new custom Response for the given http.Response
func newResponse(r *http.Response) *Response {
	return &Response{Response: r}
}

// ListOptions represents the options available for making listing requests.
type ListOptions struct {
	// Level of detail returned from request
	// Values: "minimal", "partial", "complete"
	Depth string `url:"depth,omitempty"`
	// Number of entities to return
	Count int `url:"count,omitempty"`
	// The page count of entities to return, Starting at 1
	Page int `url:"page,omitempty"`
	// A term for searching through entities
	Search string `url:"search,omitempty"`
	// The property on which to sort the returned data
	Sort string `url:"sort,omitempty"`
	// The direction of the applied sort
	SortDir string `url:"dir,omitempty"`
	// The field on which to order results
	OrderBy string `url:"orderBy,omitempty"`
	// A minimum last updated timestamp
	LastUpdatedAt int `url:"lastUpdatedAt,omitempty"`
}

// RestRequest provides a generic way to make a request to the Eloqua API.
// It's very general but simple performs much of the boilerplate request actions such
// as setting the correct api url and adding auth headers.
func (c *Client) RestRequest(endpoint string, method string, jsonData string) (*Response, error) {
	url := c.BaseURL
	endpoint = strings.Trim(endpoint, " /")

	// Assume rest 2.0 if not in request
	if strings.Index(endpoint, "api/") == -1 {
		url += "/api/rest/2.0/"
	} else {
		url += "/"
	}

	url += endpoint

	// fmt.Println(jsonData)
	jsonStr := []byte(jsonData)
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.authHeader)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	return newResponse(resp), err
}

// Performs a GET request and decodes the response into the provided interface
func (c *Client) getRequestDecode(endpoint string, v interface{}) (*Response, error) {
	resp, err := c.RestRequest(endpoint, "GET", "")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return resp, err
	}

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
	}

	return resp, err
}

// Performs a GET request for a listing endpoint and decodes the response into the provided interface
func (c *Client) getRequestListDecode(endpoint string, v interface{}, opts *ListOptions) (*Response, error) {

	// Create our options if not set
	if opts == nil {
		opts = &ListOptions{}
	}
	// Set a default minimal depth
	if opts.Depth == "" {
		opts.Depth = "minimal"
	}

	encoder, _ := query.Values(opts)
	endpoint += "?" + encoder.Encode()

	resp, err := c.RestRequest(endpoint, "GET", "")

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return resp, err
	}

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	// Decode response
	if resp != nil {
		err = json.NewDecoder(resp.Body).Decode(resp)
	}

	// Decode elements onto model
	if v != nil {
		err = json.Unmarshal(resp.Elements, v)
	}

	return resp, err
}

// Performs a HTTP request using the given method
// and decodes the response into the provided interface
func (c *Client) RequestDecode(endpoint string, method string, v interface{}) (*Response, error) {

	postBody := ""

	if v != nil {
		jsonString, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		postBody = string(jsonString)
	}

	resp, err := c.RestRequest(endpoint, strings.ToUpper(method), postBody)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return resp, err
	}

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
	}

	return resp, err
}

// Performs a POST request and decodes the response into the provided interface
func (c *Client) postRequestDecode(endpoint string, v interface{}) (*Response, error) {
	return c.RequestDecode(endpoint, "POST", v)
}

// Performs a PUT request and decodes the response into the provided interface
func (c *Client) putRequestDecode(endpoint string, v interface{}) (*Response, error) {
	return c.RequestDecode(endpoint, "PUT", v)
}

// Performs a DELETE request to the provided endpoint, sending the provided interface data.
func (c *Client) deleteRequest(endpoint string, v interface{}) (*Response, error) {
	postBody := ""

	if v != nil {
		jsonString, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		postBody = string(jsonString)
	}

	resp, err := c.RestRequest(endpoint, "DELETE", postBody)
	err = checkResponse(resp)

	return resp, err
}

// errorMessages lists the common meanings for each common HTTP status code.
// These are taken directly from the Eloqua documentation.
var errorMessages = map[int]string{
	301: "Login required",
	304: "Not Modified",
	400: "Bad Request",
	// 400 Alternatives:
	// There was a missing reference
	// There was a parsing error
	// There was a serialization error
	// There was a validation error
	401: "You are not authorized to make this request",
	// 401 Alternatives:
	// Login required
	// Unauthorized
	403: "Forbidden",
	// 403 Alternatives:
	// This service has not been enabled for your site
	// XSRF Protection Failure
	404: "The requested resource was not found",
	409: "There was a conflict",
	412: "The resource you are attempting to delete has dependencies, and cannot be deleted",
	413: "Storage space exceeded",
	429: "Too Many Requests",
	500: "The service has encountered an error",
	// 500 Alternatives:
	// Internal Server Error
	502: "Bad Gateway",
	503: "Service Unavailable",
	// 503 Alternatives:
	// There was a timeout processing the request
}

// checkResponse checks the Eloqua response for errors
// and returns them in a descriptive way if possible.
func checkResponse(r *Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	content, err := ioutil.ReadAll(r.Body)
	if err == nil {
		r.ErrorContent = string(content)
	}

	if message, ok := errorMessages[r.StatusCode]; ok {
		return errors.New(message)
	}

	return errors.New("There was an issue performing your request")
}

// HTMLContent represents the htmlContent model of an Eloqua email or landing page object
type HTMLContent struct {
	Type          string `json:"type,omitempty"`
	ContentSource string `json:"contentSource,omitempty"`
	HTML          string `json:"html,omitempty"`
}

// FieldValue represents the structure in which custom field values are passed
// via the API.
type FieldValue struct {
	Type  string `json:"type,omitempty"`
	ID    int    `json:"id,omitempty,string"`
	Value string `json:"value,omitempty"`
}

// Hyperlink is an Eloqua hyperlink object that is commonly
// contained in Eloqua assets such as emails and landing pages.
type Hyperlink struct {
	Type string `json:"type,omitempty"`
	ID   int    `json:"id,omitempty,string"`
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}

// FieldMerge is an Eloqua FieldMerge Object.
// The fields available depend on the merge source (or type).
type FieldMerge struct {
	Type                  string                `json:"type,omitempty"`
	ID                    int                   `json:"id,omitempty,string"`
	Depth                 string                `json:"depth,omitempty"`
	Name                  string                `json:"name,omitempty"`
	FolderID              int                   `json:"folderId,omitempty,string"`
	Syntax                string                `json:"syntax,omitempty"`
	UpdatedAt             int                   `json:"updatedAt,omitempty,string"`
	UpdatedBy             int                   `json:"updatedBy,omitempty,string"`
	ContactFieldID        int                   `json:"contactFieldId,omitempty,string"`
	DefaultValue          string                `json:"defaultValue,omitempty"`
	MergeType             string                `json:"mergeType,omitempty"`
	DefaultContentSection DynamicContentSection `json:"defaultContentSection,omitempty"`
}

// DynamicContent represents Eloqua Dynamic Content objects.
// These can be found as part of emails or landing pages.
type DynamicContent struct {
	Type        string               `json:"type,omitempty"`
	ID          int                  `json:"id,omitempty,string"`
	Depth       string               `json:"depth,omitempty"`
	Name        string               `json:"name,omitempty"`
	FolderID    int                  `json:"folderId,omitempty,string"`
	UpdatedAt   int                  `json:"updatedAt,omitempty,string"`
	UpdatedBy   int                  `json:"updatedBy,omitempty,string"`
	CreatedAt   int                  `json:"createdAt,omitempty,string"`
	CreatedBy   int                  `json:"createdBy,omitempty,string"`
	Permissions []string             `json:"permissions,omitempty"`
	Rules       []DynamicContentRule `json:"rules,omitempty"`
}

// DynamicContentSection represents the 'section' of content of an Eloqua dynmaic content object.
type DynamicContentSection struct {
	Type        string      `json:"type,omitempty"`
	ID          int         `json:"id,omitempty,string"`
	Depth       string      `json:"depth,omitempty"`
	Name        string      `json:"name,omitempty"`
	FolderID    int         `json:"folderId,omitempty,string"`
	Permissions []string    `json:"permissions,omitempty"`
	UpdatedAt   int         `json:"updatedAt,omitempty,string"`
	UpdatedBy   int         `json:"updatedBy,omitempty,string"`
	CreatedAt   int         `json:"createdAt,omitempty,string"`
	CreatedBy   int         `json:"createdBy,omitempty,string"`
	ContentHTML string      `json:"contentHtml,omitempty"`
	ContentText string      `json:"contentText,omitempty"`
	Forms       []Form      `json:"forms,omitempty"`
	Hyperlinks  []Hyperlink `json:"hyperlinks,omitempty"`
	Images      []Image     `json:"images,omitempty"`
	Scope       string      `json:"scope,omitempty"`
	Size        Size        `json:"size,omitempty"`
}

// DynamicContentRule specifies the conditions for Dynamic Content to be shown.
type DynamicContentRule struct {
	Type           string                  `json:"type,omitempty"`
	ID             int                     `json:"id,omitempty,string"`
	ContentSection DynamicContentSection   `json:"contentSection,omitempty"`
	Depth          string                  `json:"depth,omitempty"`
	Statement      int                     `json:"statement,omitempty,string"`
	Criteria       []ContactFieldCriterion `json:"criteria,omitempty"`
}

// ContactFieldCriterion is a simple criteria for finding a match against the contact record.
type ContactFieldCriterion struct {
	Type      string             `json:"type,omitempty"`
	ID        int                `json:"id,omitempty,string"`
	FieldID   int                `json:"fieldId,omitempty,string"`
	Condition TextValueCondition `json:"condition,omitempty"`
}

// TextValueCondition represents the conditional matching of simple text values.
type TextValueCondition struct {
	Type     string `json:"type,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

// Size is a universal Eloqua object to simply track width & height
// of other assets such as images.
type Size struct {
	Type   string `json:"type,omitempty"`
	Width  int    `json:"width,omitempty,string"`
	Height int    `json:"height,omitempty,string"`
}

// TypeObject is a very simple model of an eloqua object containing
// only the 'type' field. This is used in cases where there are no extra field,
// Just a type available.
type TypeObject struct {
	Type string `json:"type,omitempty"`
}

// Position simply stores x/y coordinates. It's used for the positioning of other elements
// such as campaign elements.
type Position struct {
	Type string `json:"type,omitempty"`
	X    int    `json:"x,omitempty,string"`
	Y    int    `json:"y,omitempty,string"`
}
