// http://golang.org/pkg/strings/
// http://stackoverflow.com/questions/16551354/how-to-split-string-and-assign
// http://stackoverflow.com/questions/22688010/how-to-trim-leading-and-trailing-white-spaces-of-a-string-in-golang
package locale

import (
	"strings"
	"strconv"
)

type LangQ struct {
	Lang	string
	Q	float64
}

func ParseAcceptLanguage(acptLang string) []LangQ {
	var lqs []LangQ

	langQStrs := strings.Split(acptLang, ",")
	for _, langQStr := range langQStrs {
		trimedLangQStr := strings.Trim(langQStr, " ")

		langQ := strings.Split(trimedLangQStr, ";")
		if len(langQ) == 1 {
			lq := LangQ{langQ[0], 1}
			lqs = append(lqs, lq)
		} else {
			qp := strings.Split(langQ[1], "=")
			q, err := strconv.ParseFloat(qp[1], 64)
			if err != nil { panic(err) }
			lq := LangQ{langQ[0], q}
			lqs = append(lqs, lq)
		}
	}
	return lqs
}
