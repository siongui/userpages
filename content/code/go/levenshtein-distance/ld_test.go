package ld

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	if LevenshteinDistance("soccer", "otter") != 3 {
		t.Error(`"soccer", "otter"`)
	}

	if LevenshteinDistance("你好他他是誰", "好我我是誰") != 3 {
		t.Error(`"你好他他是誰", "好我我是誰"`)
	}

	if LevenshteinDistance("kitten", "sitten") != 1 {
		t.Error(`"kitten", "sitten"`)
	}

	if LevenshteinDistance("sitten", "sittin") != 1 {
		t.Error(`"sitten", "sittin"`)
	}

	if LevenshteinDistance("sittin", "sitting") != 1 {
		t.Error(`"sittin", "sitting"`)
	}

	if LevenshteinDistance("sort", "some") != 2 {
		t.Error(`"sort", "some"`)
	}

	if LevenshteinDistance("sort", "same") != 3 {
		t.Error(`"sort", "same"`)
	}

	if LevenshteinDistance("sort", "soft") != 1 {
		t.Error(`"sort", "soft"`)
	}
}
