package readpdftext

import (
	"fmt"
)

func ExampleReadPlainTextFromPDF() {
	content, err := ReadPlainTextFromPDF("test.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}
