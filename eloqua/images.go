package eloqua

import (
	"fmt"
)

// ImageService provides access to all the endpoints related
// to image data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Images/images-API.htm
type ImageService struct {
	client *Client
}

// Image represents an Eloqua image object.
// Images are often in other Eloqua models such as Emails & Landing Pages
type Image struct {
	Type      string `json:"type,omitempty"`
	ID        int    `json:"id,omitempty,string"`
	CreatedAt int    `json:"createdAt,omitempty,string"`
	CreatedBy int    `json:"createdBy,omitempty,string"`
	Depth     string `json:"depth,omitempty"`
	Name      string `json:"name,omitempty"`
	FolderID  int    `json:"folderId,omitempty,string"`

	UpdatedAt    int      `json:"updatedAt,omitempty,string"`
	UpdatedBy    int      `json:"updatedBy,omitempty,string"`
	Permissions  []string `json:"permissions,omitempty"`
	FullImageURL string   `json:"fullImageUrl,omitempty"`
	Size         Size     `json:"size,omitempty"`
	ThumbnailURL string   `json:"thumbnailUrl,omitempty"`
}

// Create a new image in eloqua
func (e *ImageService) Create(name string, image *Image) (*Image, *Response, error) {
	if image == nil {
		image = &Image{}
	}
	image.Name = name

	endpoint := "/assets/image"
	resp, err := e.client.postRequestDecode(endpoint, image)
	return image, resp, err
}

// Get an image object via its ID
func (e *ImageService) Get(id int) (*Image, *Response, error) {
	endpoint := fmt.Sprintf("/assets/image/%d?depth=complete", id)
	image := &Image{}
	resp, err := e.client.getRequestDecode(endpoint, image)
	return image, resp, err
}

// List many eloqua images
func (e *ImageService) List(opts *ListOptions) ([]Image, *Response, error) {
	endpoint := "/assets/images"
	images := new([]Image)
	resp, err := e.client.getRequestListDecode(endpoint, images, opts)
	return *images, resp, err
}

// Update an existing image in eloqua
func (e *ImageService) Update(id int, name string, image *Image) (*Image, *Response, error) {
	if image == nil {
		image = &Image{}
	}

	image.ID = id
	image.Name = name

	endpoint := fmt.Sprintf("/assets/image/%d", image.ID)
	resp, err := e.client.putRequestDecode(endpoint, image)
	return image, resp, err
}

// Delete an existing image from eloqua
func (e *ImageService) Delete(id int) (*Response, error) {
	image := &Image{ID: id}
	endpoint := fmt.Sprintf("/assets/image/%d", image.ID)
	resp, err := e.client.deleteRequest(endpoint, image)
	return resp, err
}
