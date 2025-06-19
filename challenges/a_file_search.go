package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type FileFilter func(d fs.DirEntry, fullPath string) bool

func compose(filters ...FileFilter) FileFilter {
	return func(d fs.DirEntry, p string) bool {
		for _, f := range filters {
			if !f(d, p) {
				return false
			}
		}
		return true
	}
}

func NameContains(substr string) FileFilter {
	return func(d fs.DirEntry, _ string) bool {
		return strings.Contains(d.Name(), substr)
	}
}

func Hidden(hidden bool) FileFilter {
	return func(d fs.DirEntry, _ string) bool {
		if hidden {
			return true
		}

		isHidden := false

		if strings.HasPrefix(d.Name(), ".") {
			isHidden = true
		}

		return !isHidden
	}
}

func ListFiles(path string, f FileFilter) ([]string, error) {
	var out []string

	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Type().IsRegular() && f(d, p) {
			out = append(out, p)
		}
		return nil
	})

	return out, err
}

func main() {
	filtro := compose(
		Hidden(true),
		NameContains("list"),
	)

	files, err := ListFiles("../linked_lists", filtro)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	for _, f := range files {
		fmt.Println(f)
	}
}
