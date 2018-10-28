package main

import (
	"fmt"
	"strconv"
)

const MaxDigits = 201
const BASE = 10

func MakePositiveInt(s string) (n [MaxDigits]int) {
	// make n zero
	for i := 0; i < MaxDigits; i++ {
		n[i] = 0
	}

	for index, digit := range s {
		i := len(s) - index - 1
		switch digit {
		case '0':
			n[i] = 0
		case '1':
			n[i] = 1
		case '2':
			n[i] = 2
		case '3':
			n[i] = 3
		case '4':
			n[i] = 4
		case '5':
			n[i] = 5
		case '6':
			n[i] = 6
		case '7':
			n[i] = 7
		case '8':
			n[i] = 8
		case '9':
			n[i] = 9
		default:
			panic("invalid digit in number string")
		}
	}

	return
}

func AddPositiveInt(a, b [MaxDigits]int) (c [MaxDigits]int) {
	var carry, sum = 0, 0

	// make c zero
	for i := 0; i < MaxDigits; i++ {
		c[i] = 0
	}

	for i := 0; i < MaxDigits; i++ {
		sum = a[i] + b[i] + carry

		if sum >= BASE {
			carry = 1
			sum -= BASE
		} else {
			carry = 0
		}

		c[i] = sum
	}

	if carry != 0 {
		panic("overflow in addition")
	}

	return
}

func MultiplyOneDigit(a [MaxDigits]int, n int) (b [MaxDigits]int) {
	var carry = 0

	// make b zero
	for i := 0; i < MaxDigits; i++ {
		b[i] = 0
	}

	for i := 0; i < MaxDigits; i++ {
		b[i] = n * a[i]

		b[i] += carry

		if b[i] >= BASE {
			carry = b[i] / BASE
			b[i] %= BASE
		} else {
			carry = 0
		}
	}

	if carry != 0 {
		panic("overflow in multiplication")
	}

	return
}

func ShiftLeft(a [MaxDigits]int, n int) [MaxDigits]int {
	var i int

	for i = MaxDigits - 1; i >= n; i-- {
		a[i] = a[i-n]
	}
	for i >= 0 {
		a[i] = 0
		i -= 1
	}

	return a
}

func MultiplyPositiveInt(a, b [MaxDigits]int) (c [MaxDigits]int) {
	// make c zero
	for i := 0; i < MaxDigits; i++ {
		c[i] = 0
	}

	for i := 0; i < MaxDigits; i++ {
		tmp := MultiplyOneDigit(b, a[i])
		tmp = ShiftLeft(tmp, i)
		c = AddPositiveInt(c, tmp)
	}

	return
}

// a^b, b is integer and b>=2
func Power(a [MaxDigits]int, b int) (c [MaxDigits]int) {
	c = MultiplyPositiveInt(a, a)
	for i := 2; i < b; i++ {
		c = MultiplyPositiveInt(c, a)
	}
	return
}

func PositiveIntToString(a [MaxDigits]int) (s string) {
	isLeadingZero := true
	for i := MaxDigits - 1; i >= 0; i-- {
		if isLeadingZero && a[i] == 0 {
			continue
		} else {
			isLeadingZero = false
			s += strconv.Itoa(a[i])
		}
	}
	return
}

func main() {
	distinctNum := make(map[string]bool)
	for i := 2; i <= 100; i++ {
		a := MakePositiveInt(strconv.Itoa(i))
		for b := 2; b <= 100; b++ {
			c := Power(a, b)
			distinctNum[PositiveIntToString(c)] = true
		}
	}

	fmt.Println(len(distinctNum))
}
