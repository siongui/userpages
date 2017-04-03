package ld

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func minimum(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return min
}

// lenS and lenT are the number of characters in string s and t respectively
func Recursive(s []rune, lenS int, t []rune, lenT int) int {
	var cost int

	// base case: empty strings
	if lenS == 0 {
		return lenT
	}
	if lenT == 0 {
		return lenS
	}

	// test if last characters of the strings match
	if s[lenS-1] == t[lenT-1] {
		cost = 0
	} else {
		cost = 1
	}

	// return minimum of delete char from s, delete char from t, and delete char from both
	return minimum(
		Recursive(s, lenS-1, t, lenT)+1,
		Recursive(s, lenS, t, lenT-1)+1,
		Recursive(s, lenS-1, t, lenT-1)+cost,
	)
}

func LevenshteinDistance(s, t string) int {
	srs := StringToRuneSlice(s)
	trs := StringToRuneSlice(t)
	return Recursive(srs, len(srs), trs, len(trs))
}
