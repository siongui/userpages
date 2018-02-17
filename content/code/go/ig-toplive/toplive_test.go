package instago

import (
	"os"
	"testing"
)

func ExampleToplive(t *testing.T) {
	tl, err := Toplive(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tl)
}
