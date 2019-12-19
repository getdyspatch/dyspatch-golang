package dyspatch_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/getdyspatch/dyspatch-golang"
	"golang.org/x/net/context"
)

const version = "application/vnd.dyspatch.2019.10+json"

func bodyToString(body io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}

func TestMain(t *testing.T) {
	cfg := dyspatch.NewConfiguration()
	auth := context.WithValue(context.Background(), dyspatch.ContextAPIKey, dyspatch.APIKey{
		Key:    os.Getenv("DYSPATCH_API_KEY"),
		Prefix: "Bearer",
	})

	client := dyspatch.NewAPIClient(cfg)
	client.TemplatesApi.GetTemplates(auth, version, nil)
	templates, _, err := client.TemplatesApi.GetTemplates(auth, version, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Print out Templates
	for _, template := range templates.Data {
		fmt.Println(template.Name)
	}

	// Get Drafts
	drafts, _, err := client.DraftsApi.GetDrafts(auth, version, nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", drafts)
	for _, draft := range drafts.Data {
		fmt.Printf("%v\n", draft)
	}

	// Page through drafts
	cursorVal := ""
	keepPaging := true
	for {
		drafts, _, _ := client.DraftsApi.GetDrafts(auth, version, &dyspatch.GetDraftsOpts{
			Cursor: optional.NewString(cursorVal),
		})
		cursorVal = drafts.Cursor.Next
		keepPaging = drafts.Cursor.HasMore

		fmt.Printf("%v\n", drafts.Cursor.Next)
		if !keepPaging {
			break
		}
	}
}
