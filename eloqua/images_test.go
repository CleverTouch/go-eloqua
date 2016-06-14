package eloqua

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestImageCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &Image{Name: "A Test Image"}

	addRestHandlerFunc("/assets/image", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "POST")
		v := new(Image)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Image.Create body", v, input)

		fmt.Fprint(w, `{"type":"Image","id":"10005","name":"A Test Image"}`)
	})

	image, _, err := client.Images.Create("A Test Image", nil)
	if err != nil {
		t.Errorf("Images.Create recieved error: %v", err)
	}

	output := &Image{ID: 10005, Name: "A Test Image", Type: "Image"}

	testModels(t, "Images.Create", image, output)
}

func TestImageGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/image/1005", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "complete")
		testMethod(t, req, "GET")
		fmt.Fprint(w, `{"assetType":"Image","id":"10005","name":"A Test Image", "folderId": "101","size":{"type":"Size","width":"650","height":"80"}}`)
	})

	image, _, err := client.Images.Get(1005)
	if err != nil {
		t.Errorf("Images.Get recieved error: %v", err)
	}

	output := &Image{ID: 10005, Name: "A Test Image", FolderID: 101, Size: Size{Type: "Size", Width: 650, Height: 80}}
	testModels(t, "Images.Get", image, output)
}

func TestImageList(t *testing.T) {
	setup()
	defer teardown()

	reqOpts := &ListOptions{Count: 100, Page: 1}

	addRestHandlerFunc("/assets/images", func(w http.ResponseWriter, req *http.Request) {
		testURLParam(t, req, "depth", "minimal")
		testURLParam(t, req, "count", "100")
		testURLParam(t, req, "page", "1")
		testMethod(t, req, "GET")

		rJSON := `{"elements":[{"type":"Image","id":"10005","name":"A Test Image"}], "page":1,"pageSize":100,"total":1}`
		fmt.Fprint(w, rJSON)
	})

	images, resp, err := client.Images.List(reqOpts)
	if err != nil {
		t.Errorf("Images.List recieved error: %v", err)
	}

	want := []Image{{Type: "Image", ID: 10005, Name: "A Test Image"}}
	testModels(t, "Images.List", images, want)

	if resp.PageSize != reqOpts.Count {
		t.Error("Images.List response page size incorrect")
	}
	if resp.Page != reqOpts.Page {
		t.Error("Images.List response page number incorrect")
	}
}

func TestImageUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Image{ID: 10005, Name: "Updated Image", FolderID: 10}

	addRestHandlerFunc("/assets/image/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Image)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Images.Update body", v, input)

		fmt.Fprintf(w, `{"type":"Image","id":"10005","name":"%s","folderId":"10"}`, v.Name)
	})

	image, _, err := client.Images.Update(10005, "Updated Image", input)
	if err != nil {
		t.Errorf("Images.Update recieved error: %v", err)
	}

	testModels(t, "Images.Update", image, input)
}

func TestImageUpdateWithoutPassingModel(t *testing.T) {
	setup()
	defer teardown()

	input := &Image{ID: 10005, Name: "Updated Image"}

	addRestHandlerFunc("/assets/image/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "PUT")
		v := new(Image)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Images.Update(Without Model) body", v, input)

		fmt.Fprintf(w, `{"type":"Image","id":"10005","name":"%s","fullImageUrl":"http://placehold.it/400x200"}`, v.Name)
	})

	image, _, err := client.Images.Update(10005, "Updated Image", nil)
	if err != nil {
		t.Errorf("Images.Update(Without Model) recieved error: %v", err)
	}

	input.FullImageURL = "http://placehold.it/400x200"
	input.Type = "Image"

	testModels(t, "Images.Update(Without Model)", image, input)
}

func TestImageDelete(t *testing.T) {
	setup()
	defer teardown()

	input := &Image{ID: 10005}

	addRestHandlerFunc("/assets/image/10005", func(w http.ResponseWriter, req *http.Request) {
		testMethod(t, req, "DELETE")
		v := new(Image)
		json.NewDecoder(req.Body).Decode(v)
		testModels(t, "Images.Delete body", v, input)
		w.WriteHeader(200)
	})

	resp, err := client.Images.Delete(10005)
	if err != nil {
		t.Errorf("Images.Delete recieved error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Images.Delete request failed")
	}
}
