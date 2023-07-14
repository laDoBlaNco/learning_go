package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
	sig, err = sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

/*
if file name ends with .gz

	$ cat http.log.gz | gunzip | sha1sum

else

	$ cat http.log.gz | sha1sum
*/
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName) //returning a pointer to an os.File type and an err
	if err != nil {
		return "", err
	}
	// In go we don't do the Try/Finally/Close, we use defer which will always happen
	// after the function exits. So we close  the file when we are done and don't worry about it
	defer file.Close() // defer is always at the function level. If you have several defers
	// they will be call in reverse order. They are in a defer stack LIFO order
	// idomatic: acquire a resource, check for error, defer release
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file) // since file type implements a Read method, its a 'Reader'
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}

	//io.CopyN(os.Stdout, r, 100)// up to now I'm getting the  uncompressed data read
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x\n", sig), nil // %x returns the hex number
}
