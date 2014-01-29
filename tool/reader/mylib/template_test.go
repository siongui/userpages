package mylib

import (
	"testing"
	"os"
)


func TestGetTemplate(t *testing.T) {
	const filepath = "../view/index.html"
	tmpl := GetTemplate(filepath)
	tmpl.Execute(os.Stdout, nil)
}
