package readpdftext

import (
	"bytes"

	"github.com/ledongthuc/pdf"
)

func ReadPlainTextFromPDF(pdfpath string) (text string, err error) {
	f, r, err := pdf.Open(pdfpath)
	defer f.Close()
	if err != nil {
		return
	}

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return
	}

	buf.ReadFrom(b)
	text = buf.String()
	return
}
