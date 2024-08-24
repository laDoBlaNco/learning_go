package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"errors" 
)

var pl = fmt.Println

func main() {
	// almost all io operations return 2 values cuz there's always a chance of
	// something failing.
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// the above is an idiomatic Go patter for using any resources que sea files,
	// connections, Dbase, sockets, etc (open connection, check for error,
	// defer the close) siempre

	primeNums := []int{2, 3, 5, 7, 11}
	var strPrimeNums []string
	for _, i := range primeNums {
		strPrimeNums = append(strPrimeNums, strconv.Itoa(i))
	}
	for _, num := range strPrimeNums {
		_, err := f.WriteString(num + "\n")
		// here we are just catching the error if there is one, but the actual
		// result of this function isn't something we are going to use. It probably
		// returns the number of bytes written as it looks like a io.Writer interface
		if err != nil {
			log.Fatal(err)
		}

	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime:", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
	
	// if we want to append the file:
	pl()
	_,err = os.Stat("data.txt") 
	if errors.Is(err,os.ErrNotExist){
		pl("File doesn't exist") 
	}
	f,err = os.OpenFile("data.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		log.Fatal(err) 
	}
	defer f.Close() 
	if _,err:=f.WriteString("13\n");err!=nil{
		log.Fatal(err) 
	}

}
