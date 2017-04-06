package ld

import (
	"encoding/json"
	"os"
	"testing"
)

func ExampleBKTree(t *testing.T) {
	bkt := NewBKTree()
	bkt.Add("some")
	bkt.Add("soft")
	bkt.Add("same")
	bkt.Add("mole")
	bkt.Add("soda")
	bkt.Add("salmon")

	b, err := json.MarshalIndent(bkt.root, "", "  ")
	if err != nil {
		t.Error(err)
	}
	os.Stdout.Write(b)

	r := bkt.Search("sort", 2)
	for _, node := range r {
		println(node.Str)
	}
}
