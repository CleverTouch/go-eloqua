package eloqua

// VisitorService provides access to all the endpoints related
// to Visitor data within eloqua
//
// Eloqua API docs: https://goo.gl/r5Ctvr
type VisitorService struct {
	client *Client
}

// Visitor represents an Eloqua Visitor object.
type Visitor struct {
	Type                 string `json:"type,omitempty"`
	VisitorID            int    `json:"visitorId,omitempty,string"`
	CreatedAt            int    `json:"createdAt,omitempty,string"`
	IPAddress            string `json:"V_IPAddress,omitempty"`
	LastVisitDateAndTime int    `json:"V_LastVisitDateAndTime,omitempty,string"`
	ExternalID           string `json:"externalId,omitempty"`
	ContactID            int    `json:"contactId,omitempty,string"`
	CurrentStatus        string `json:"currentStatus,omitempty"`
}

// List many eloqua visitors
func (e *VisitorService) List(opts *ListOptions) ([]Visitor, *Response, error) {
	endpoint := "/data/visitors"
	visitors := new([]Visitor)
	resp, err := e.client.getRequestListDecode(endpoint, visitors, opts)
	return *visitors, resp, err
}
