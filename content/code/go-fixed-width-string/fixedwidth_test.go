package fixedwidth

import "testing"

func TestFixedLengthString(t *testing.T) {
	println(FixedLengthString(20, "This is"))
	println(FixedLengthString(20, "an example of"))
	println(FixedLengthString(20, "fixed width/length"))
	println(FixedLengthString(20, "string"))
}
