package rstmeta

import (
	"os"
	"testing"
)

func TestRstMeta(t *testing.T) {
	writeRstMetadata(os.Stdout, "test title 測試標題", "2016-04-22T20:58+08:00", "tag1, tag2", "category1", "my summary")
}
