package converter

import (
	"testing"
)

func TestKebabCaseToCamelCase(t *testing.T) {
	kebabString := "border-left-color"
	t.Log(kebabString)
	camel := kebabToCamelCase(kebabString)
	if camel != "borderLeftColor" {
		t.Error(camel)
		return
	}
	t.Log(camel)
}
