package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// ...for quick debugging
var p = fmt.Println

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// create a new sub-dir in the current directory
	err := os.Mkdir("subdir", 0755)
	check(err)
	// when we create temp directories its a good practice to defer their
	// removal os.RemoveAll will delete a whole dir tree
	defer os.RemoveAll("subdir")

	// This helper function will create a new empty file.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// We can create a hierarchy of dirs, including parents with MkdirAll.
	// this is similar to mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/file2")
	createEmptyFile("subdir/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir list directory contents, returning a slice of os.DirEntry objs
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir lets us change the current working directory, similar to 'cd'
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Now we'll see the contents of subdir/parent/child when listing
	// the current dir.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to where we started
	err = os.Chdir("../../..")
	check(err)

	// We can also visit a directory recursively, including all its sub-directories
	// Walk accepts a callback function to handle every file or directory
	// visited.
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)

}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
