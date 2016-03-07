package rmcmt

import "testing"

var ccode = []byte(`
/*
 * I am C style comments
 */
int main() {
  return 0;
}
`)

var cppcode = []byte(`
// I am Cpp style comments
int main() {
  return 0;
}
`)

var codeAfter = []byte(`

int main() {
  return 0;
}
`)

func TestRemoveCStyleComments(t *testing.T) {
	if string(RemoveCStyleComments(ccode)) != string(codeAfter) {
		t.Error("Remove C style comments failure!")
	}
}

func TestRemoveCppStyleComments(t *testing.T) {
	if string(RemoveCppStyleComments(cppcode)) != string(codeAfter) {
		t.Error("Remove Cpp style comments failure!")
	}
}
