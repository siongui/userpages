package igpost

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetAllPostCode(t *testing.T) {
	codes, err := GetAllPostCode(
		os.Getenv("IG_TEST_USERNAME"),
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		t.Error(err)
		return
	}
	for _, code := range codes {
		fmt.Printf("URL: https://www.instagram.com/p/%s/\n", code)
	}
}
