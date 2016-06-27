package eloqua

import (
	"fmt"
)

// OptionListService provides access to all the endpoints related
// to Option List data within eloqua
//
// Eloqua API docs: https://goo.gl/KyNNO1
type OptionListService struct {
	client *Client
}

// OptionList represents an Eloqua Option List object.
// OptionLists are also known as picklists or select lists
type OptionList struct {
	Type        string   `json:"type,omitempty"`
	ID          int      `json:"id,omitempty,string"`
	Depth       string   `json:"depth,omitempty"`
	Name        string   `json:"name,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Elements    []Option `json:"elements,omitempty"`
}

// Option represents an Eloqua select/picklist option.
type Option struct {
	Type        string `json:"type,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
	Value       string `json:"value,omitempty"`
}

// Create a new optionList in eloqua
func (e *OptionListService) Create(name string, optionList *OptionList) (*OptionList, *Response, error) {
	if optionList == nil {
		optionList = &OptionList{}
	}
	optionList.Name = name

	endpoint := "/assets/optionList"
	resp, err := e.client.postRequestDecode(endpoint, optionList)
	return optionList, resp, err
}

// Get an optionList object via its ID
func (e *OptionListService) Get(id int) (*OptionList, *Response, error) {
	endpoint := fmt.Sprintf("/assets/optionList/%d?depth=complete", id)
	optionList := &OptionList{}
	resp, err := e.client.getRequestDecode(endpoint, optionList)
	return optionList, resp, err
}

// List many eloqua optionLists
func (e *OptionListService) List(opts *ListOptions) ([]OptionList, *Response, error) {
	endpoint := "/assets/optionLists"
	optionLists := new([]OptionList)
	resp, err := e.client.getRequestListDecode(endpoint, optionLists, opts)
	return *optionLists, resp, err
}

// Update an existing optionList in eloqua
// Updating will delete all current options.
func (e *OptionListService) Update(id int, name string, optionList *OptionList) (*OptionList, *Response, error) {
	if optionList == nil {
		optionList = &OptionList{}
	}

	optionList.ID = id
	optionList.Name = name

	endpoint := fmt.Sprintf("/assets/optionList/%d", optionList.ID)
	resp, err := e.client.putRequestDecode(endpoint, optionList)
	return optionList, resp, err
}

// Delete an existing optionList from eloqua
func (e *OptionListService) Delete(id int) (*Response, error) {
	optionList := &OptionList{ID: id}
	endpoint := fmt.Sprintf("/assets/optionList/%d", optionList.ID)
	resp, err := e.client.deleteRequest(endpoint, optionList)
	return resp, err
}
