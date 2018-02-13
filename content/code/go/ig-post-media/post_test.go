package igmedia

import (
	"os"
	"testing"
)

func ExampleGetPostInfo(t *testing.T) {
	em, err := GetPostInfo(os.Getenv("IG_TEST_CODE"),
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		t.Error(err)
		return
	}
	em.printEdgeMediaInfo()
}
