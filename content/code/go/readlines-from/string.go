package lines

import "strings"

func StringToLines(s string) ([]string, error) {
	r := strings.NewReader(s)
	return LinesFromReader(r)
}
