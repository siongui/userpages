package ld

import (
	"testing"
)

func TestEditDistance(t *testing.T) {
	if EditDistance("soccer", "otter") != 3 {
		t.Error(`"soccer", "otter"`)
	}

	if EditDistance("你好他他是誰", "好我我是誰") != 3 {
		t.Error(`"你好他他是誰", "好我我是誰"`)
	}

	if EditDistance("kitten", "sitten") != 1 {
		t.Error(`"kitten", "sitten"`)
	}

	if EditDistance("sitten", "sittin") != 1 {
		t.Error(`"sitten", "sittin"`)
	}

	if EditDistance("sittin", "sitting") != 1 {
		t.Error(`"sittin", "sitting"`)
	}

	if EditDistance("sort", "some") != 2 {
		t.Error(`"sort", "some"`)
	}

	if EditDistance("sort", "same") != 3 {
		t.Error(`"sort", "same"`)
	}

	if EditDistance("sort", "soft") != 1 {
		t.Error(`"sort", "soft"`)
	}
}
