/*
Write a function that gets an index file with names of files and sha256
signatures in the following format
0c4ccc63a912bbd6d45174251415c089522e5c0e75286794ab1f86cb8e2561fd  taxi-01.csv
f427b5880e9164ec1e6cda53aa4b2d1f1e470da973e5b51748c806ea5c57cbdf  taxi-02.csv
4e251e9e98c5cb7be8b34adfcb46cc806a4ef5ec8c95ba9aac5ff81449fc630c  taxi-03.csv
...

You should compute concurrently sha256 signatures of these files and see if
they match the ones in the index file.

  - Print the number of processed files
  - If there's a mismatch, print the offending file(s) and exit the program with
    non-zero value

Grab taxi-sha256.zip from the web site and open it. The index file is sha256sum.txt
*/
package main

import (
	"bufio"
	"compress/bzip2"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func fileSig(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, bzip2.NewReader(file))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Parse signature file. Return map of path->signature
func parseSigFile(r io.Reader) (map[string]string, error) {
	sigs := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Line example
		// 6c6427da7893932731901035edbb9214  nasa-00.log
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			// TODO: line number
			return nil, fmt.Errorf("bad line: %q", scanner.Text())
		}
		sigs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}

func main() {
	rootDir := "./data" // Change to where to unzipped taxi-sha256.zip
	file, err := os.Open(path.Join(rootDir, "sha256sum.txt"))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	sigs, err := parseSigFile(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	start := time.Now()
	ok := true
	// create a channel to get our new result struct
	ch := make(chan result)
	for name, signature := range sigs {
		fileName := path.Join(rootDir, name) + ".bz2"
		// sig, err := fileSig(fileName)
		go sigWorker(fileName, signature, ch)
		// so here we are sending up just the processing of the files to the goroutine and not
		// the everything in this for loop.
	}
	// then here we create another for loop around the gathering of the information from ch
	for range sigs {
		r := <-ch
		if r.err != nil {
			fmt.Fprintf(os.Stderr, "error: %s - %s\n", r.fileName, err)
			ok = false
			continue
		}

		if !r.match {
			ok = false
			fmt.Printf("error: %s mismatch\n", r.fileName)
		}
	}

	duration := time.Since(start)
	fmt.Printf("processed %d files in %v\n", len(sigs), duration)
	if !ok {
		os.Exit(1)
	}
}

// first we need to think about the context, thinking about what we  want to return
// from the go routine. So we create a struct or record of what is happening on each routine
type result struct {
	fileName string
	err      error
	match    bool
}

// let's create a worker to work go up wtih our goroutine and manage the work happening in that
// branch
// NOTE: If we have sequential code that works, don't touch it. We can wrap it easily in a
// goroutine without refactoring all of our code. "Leave concurrency to the user" o sea create
// code as you want it first and then add concurrency later if needed. That's the Go way.
func sigWorker(fileName, signature string, ch chan<- result) {
	r := result{fileName: fileName}
	sig, err := fileSig(fileName)
	if err != nil {
		r.err = err
	} else {
		r.match = sig == signature
	}
	ch <- r
}

// we could have put this code in the for loop as I was thinking of originally, but when the
// code gets complicated its better to put it on the side as we do here.

// Concurrency preference in Go is with channels, but if we must use sync constructs then we
// have the sync package, which we'll look at next. 

