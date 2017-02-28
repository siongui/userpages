package write

import (
	"fmt"
	"os"
)

func BadWriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = fmt.Fprintf(fo, s)
	if err != nil {
		return err
	}

	return nil
}
