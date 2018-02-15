// Sort file by size
package sortfile

import (
	"os"
	"sort"
)

func SortFileBySize(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() < files[j].Size()
	})
}
