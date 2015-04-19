package builder

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//dirInfo get the dirs & files under the root path.
type dirInfo struct {
	root  string
	dirs  []string
	files []string
}

func (di *dirInfo) String() string {
	buf := bytes.NewBufferString(fmt.Sprintf("Root:\n\t%s\n", di.root))
	buf.WriteString("\nDirs:\n")
	for _, d := range di.dirs {
		buf.WriteString("\t")
		buf.WriteString(d)
		buf.WriteString("\n")
	}

	buf.WriteString("\nFiles:\n")
	for _, f := range di.files {
		buf.WriteString("\t")
		buf.WriteString(f)
		buf.WriteString("\n")
	}

	return buf.String()

}

func newDirInfo(root string) *dirInfo {
	return &dirInfo{
		root:  filepath.Clean(root),
		dirs:  make([]string, 0),
		files: make([]string, 0),
	}

}

func (di *dirInfo) analysis() {

	filepath.Walk(di.root, func(path string, f os.FileInfo, err error) error {

		path = strings.TrimPrefix(path, di.root)
		if path == "" {
			return err
		}

		if f.IsDir() {
			di.dirs = append(di.dirs, path)
		} else {
			di.files = append(di.files, path)
		}
		return err
	})

}
