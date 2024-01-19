package freax

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ListDir(name string, recursive bool, hidden bool) ([]string, error) {

	entries, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, len(entries))

	for i := range entries {

		// Skip hidden (prefixed with a '.') entries if hidden is disabled
		if !hidden && strings.HasPrefix(entries[i].Name(), ".") {
			continue
		}

		list = append(list, entries[i].Name())

		if recursive && entries[i].IsDir() {
			e, err := ListDir(path.Join(name, entries[i].Name()), recursive, hidden)
			if err != nil {
				return list, err
			}

			for j := range e {
				list = append(list, path.Join(entries[i].Name(), e[j]))
			}
		}
	}

	return list, nil
}

func ListFiles(path string) ([]string, error) {

	list := make([]string, 0)

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {

		if !d.IsDir() {
			list = append(list, path)
		}
		return err
	})

	return list, err
}
