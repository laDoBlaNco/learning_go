// Go by Example: Temporary Files and Directories
// Throughout program execution, we often want to create data that isn't
// needed after the program exits. Temporary files and Directories are usefule
// for this purpsoe since they don't pollute the file system over time.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// The easiest way to create a temp file is by calling os.CreateTemp. It creates
// a file and opens it for readiang and writing. We provide "" as the first argument
// so os.CreateTemp will create the file in the default location for our OS
func main() {

	f, err := os.CreateTemp("", "sample")
	check(err)

	// Display the name of the temporary file. On Unix based OSes the directory will likely
	// be /tmp. The file name starts with the prefix given as the second arg to os.CreateTemp
	// and the rest is chosen automatically to ensure that concurrent calls will always create
	// different file names.
	fmt.Println("Temp file name:", f.Name())

	// You them clean up after we're done. The OS is likel to clear up temporary files
	// by itself after some time, but its a good practice to do this explicitly.
	defer os.Remove(f.Name())

	// We can then write some data into the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// If we intend to write many temporary files, we may prefer to create a temporary dir
	// os.MkdirTemp's arguments are the same as CreateTemp's, but it returns a directory
	// name rather than an open file.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Now we can synthesize temporary file names by prefixing them with our temporary
	// directory
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
