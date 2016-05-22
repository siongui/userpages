package fr

import (
	"os"
	"path/filepath"
	"regexp"
)

var redundant = regexp.MustCompile(` \(\d+\)$`)

func isRedundant(filename string) bool {
	ext := filepath.Ext(filename)
	filenameWithoutExt := filename[:len(filename)-len(ext)]
	if redundant.Match([]byte(filenameWithoutExt)) {
		return true
	}
	return false
}

func removeFile(path string) {
	println("Remove " + path)
	os.Remove(path)
}

func FindRedundant(dirPath string) {
	// walk all files in directory
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if isRedundant(info.Name()) {
				removeFile(path)
			}
		}
		return nil
	})
}
