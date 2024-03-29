package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		file1 := &File{name: "File1"}
		file2 := &File{name: "File2"}
		file3 := &File{name: "File3"}

		folder1 := &Folder{
			children: []Inode{file1},
			name:     "Folder1",
		}

		folder2 := &Folder{
			children: []Inode{folder1, file2, file3},
			name:     "Folder2",
		}
		fmt.Println("\nPrinting hierarchy for Folder2")
		folder2.print("  ")

		cloneFolder := folder2.clone()
		fmt.Println("\nPrinting hierarchy for clone Folder")
		cloneFolder.print("  ")
	})
}
