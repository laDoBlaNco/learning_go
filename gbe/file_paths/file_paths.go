// Go by Example: File Paths
// The filepath pkg provides functions to parse and construct file paths
// in a way that is portable between operating systems; dir/file on linux vs
// dir\file on Windows, for example.

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Join should be used to construct paths in a portable way. It takes
	// any number of arguments and constructs a hierarchical path from them
	p := filepath.Join("dir1", "dir2", "dir3", "filename")
	fmt.Println("p:", p)

	// You should always use join instead of concatenating /s or \s manually
	// In addition to providing portability, Join will also normalize paths by
	// removing superfluous separators and diretory changes
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// dir and base can be used to split a path to a directory and the file
	// Alternatively, Split will return both in the same call
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// We can also check whether a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Some file names have extentions following a dot. We can split the extension
	// out of such names with Ext
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// To find the file's name with the extension removed, use strings.TrimSuffix
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel files a relative path between a base and a target. It returns an error
	// if the target cannot be made relative to base
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

}

// And next in line ... Directories
