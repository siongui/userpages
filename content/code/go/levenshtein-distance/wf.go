package ld

// The edit distances here is actually Levenshtein distance, minimum number of
// single-character edits (insertions, deletions or substitutions) required to
// change one string into the other.
func EditDistance(s, t string) int {
	ss := StringToRuneSlice(s)
	m := len(ss) + 1
	tt := StringToRuneSlice(t)
	n := len(tt) + 1

	d := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		d[i][0] = i
	}

	for j := 0; j < n; j++ {
		d[0][j] = j
	}

	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			var substitutionCost int
			if ss[i-1] == tt[j-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}
			d[i][j] = minimum(
				d[i-1][j]+1,
				d[i][j-1]+1,
				d[i-1][j-1]+substitutionCost,
			)
		}
	}

	return d[m-1][n-1]
}
