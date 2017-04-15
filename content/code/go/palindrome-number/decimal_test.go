package palindrome

import (
	"testing"
)

func TestIsDecimalPalindromeNumber(t *testing.T) {
	if IsDecimalPalindromeNumber(123) != false {
		t.Error(123)
		return
	}
	if IsDecimalPalindromeNumber(1221) != true {
		t.Error(1221)
		return
	}
	if IsDecimalPalindromeNumber(-125521) != true {
		t.Error(-125521)
		return
	}
	if IsDecimalPalindromeNumber(-12521) != true {
		t.Error(-12521)
		return
	}
	if IsDecimalPalindromeNumber(12521) != true {
		t.Error(12521)
		return
	}
	if IsDecimalPalindromeNumber(-1251) != false {
		t.Error(-1251)
		return
	}
	if IsDecimalPalindromeNumber(1251) != false {
		t.Error(1251)
		return
	}
}
