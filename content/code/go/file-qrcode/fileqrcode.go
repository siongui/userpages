package fileqrcode

import (
	"encoding/base64"
	qrencode "github.com/skip2/go-qrcode"
	"github.com/tuotoo/qrcode"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// ======================
// Save file into QR code
// ======================

func filenameToBase64(filepath string) string {
	filename := path.Base(filepath)
	return base64.StdEncoding.EncodeToString([]byte(filename))
}

func FileToQrcode(filepath string) (err error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	outputFilename := filenameToBase64(filepath) + ".png"
	f64 := base64.StdEncoding.EncodeToString(b)
	err = qrencode.WriteFile(f64, qrencode.Medium, 256, outputFilename)
	return
}

// ==========================
// Retrieve file from QR code
// ==========================

func getFilenameFromBase64(encoded string) (filename string, err error) {
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return
	}
	filename = string(b)
	return
}

func getFileContentFromBase64(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

func QrcodeToFile(filepath string) (err error) {
	filename := path.Base(filepath)
	ext := path.Ext(filename)
	oriFilename, err := getFilenameFromBase64(strings.TrimSuffix(filename, ext))
	if err != nil {
		return
	}

	f, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	qrmatrix, err := qrcode.Decode(f)
	if err != nil {
		return
	}

	content, err := getFileContentFromBase64(qrmatrix.Content)
	if err != nil {
		return
	}

	return ioutil.WriteFile(oriFilename, content, 0644)
}
