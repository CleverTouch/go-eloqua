package eloqua

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestEmailGet(t *testing.T) {
	setup()
	defer teardown()

	addRestHandlerFunc("/assets/email/1", func(w http.ResponseWriter, req *http.Request) {
		testUrlParam(t, req, "depth", "complete")
		rJson := `{"type":"Email","currentStatus":"Active","id":"1","createdAt":"1444752811","createdBy":"11","depth":"complete","folderId":"10","name":"Test Email 1"}`
		io.WriteString(w, rJson)
	})

	email, _, _ := client.Emails.Get(1)
	// TODO - Look at https://github.com/google/go-github/blob/master/github/orgs_test.go#L34 to compare
	fmt.Println(email.Name)
	if email.ID != 1 || email.Name != "Test Email 1" {
		t.Errorf("Incorrect email details")
	}
}
