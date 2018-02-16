package instago

import (
	"os"
	"testing"
)

func ExampleTopsearch(t *testing.T) {
	tr, err := Topsearch("instagram",
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tr)
}
