package pfs

import (
	"fmt"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	if fmt.Sprintf("%v", PrimeFactors(23)) != `[23]` {
		t.Error(23)
	}
	if fmt.Sprintf("%v", PrimeFactors(12)) != `[2 2 3]` {
		t.Error(12)
	}
	if fmt.Sprintf("%v", PrimeFactors(360)) != `[2 2 2 3 3 5]` {
		t.Error(360)
	}
	if fmt.Sprintf("%v", PrimeFactors(97)) != `[97]` {
		t.Error(97)
	}
}
