// Trying to learn a bit more about file writing flags in os.Create cuz I'm shot
// myself in the foot twice now with my new  gopro cli

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ...for quick debugging
var p = fmt.Println

func main() {

	// so here we call os.Stat to check to see if the file exists or not
	_, err := os.Stat("words.txt")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file doesn't exist")
	} else {
		fmt.Println("file exists")
	}
	p()
	// Creating and truncating/overwriting a file
	file, err := os.Create("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("file created")
	p()

	// Deleting a file
	err = os.Remove("words.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file doesn't exist")
			return
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("file deleted")
	p()

	// Getting the file size
	// first let's create it again
	_, err = os.Create("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fInfo, err := os.Stat("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	fsize := fInfo.Size()
	fmt.Printf("The file size is %d bytes\n", fsize)
	fmt.Printf("%#v\n", fInfo)
	p()

	// Getting the last modified time
	fileName := "words.txt"
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	mTime := fileInfo.ModTime()
	fmt.Println(mTime)
	p()

	// Adding the write example here so that the read examples below have something
	// to read.
	fileName = "data.txt"
	val := `old
falcon
sky
cup
forest
`

	data := []byte(val)

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	p("done")
	p()

	// Doing the same as above a slice of strings to the file overwriting the prev.
	fileName = "data.txt"
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	words := []string{"sky", "falcon", "rock", "hawk", "old", "cup", "forest"}
	for _, word := range words {
		_, err := f.WriteString(word + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Reading a file
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	p(string(content))
	p()

	// We could also do the same with bufio to read line by line as the file could
	// be huge.
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// And now appending to a file.
	fileName = "data.txt"
	f, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString("cloud\n"); err != nil {
		log.Fatal(err)
	}

	// Now let's copy a file with ioutil
	src := "data.txt"
	dest := "data2.txt"
	
	bytesRead,err := ioutil.ReadFile(src) 
	if err!=nil{
		log.Fatal(err)
	}
	err = ioutil.WriteFile(dest,bytesRead,0644)
	if err!=nil{
		log.Fatal(err) 
	}
	p()
	p("Source copied:\n"+string(bytesRead)) 
	
	p()
	
	// Go list files:
	// here we are going to use filepath.Walk  to walk the file tree
	var files []string
	root,err:=os.Getwd()
	if err!=nil{
		log.Fatal(err)
	}
	
	err = filepath.Walk(root,func(path string,info os.FileInfo,err error)error{
		if err!=nil{
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() && filepath.Ext(path)==".txt"{
			files=append(files,path)
		}
		return nil
	})
	if err!=nil{
		log.Fatal(err)
	}
	
	for _,file:=range files{
		fmt.Println(file) 
	}
}
