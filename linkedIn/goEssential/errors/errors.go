package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {
	/*
		setupLogging()
		cfg, err := readConfig("/path/to/config.toml")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			log.Printf("error: %+v", err)
			os.Exit(1)
		}

		// normal operation
		fmt.Println(cfg)
	*/

	// Panic example
	vals := []int{1, 2, 3}

	/*
		v := vals[10] // this will get a panic
		fmt.Println(v)
	*/

	v, err := safeValue(vals, 10)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("v:", v)

	// calling panics myself
	// panic("oops!")
	fmt.Println("Don't panic")

	// CHALLENGE - Kill server
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

}

// Config holds configuration
type Config struct {
	// configuration fields go here
}

func readConfig(path string) (*Config, error) { //zero default of a pointer is nil
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}
	defer file.Close() // idiom: Open resource -> test for err -> defer close of resource

	cfg := &Config{}
	// Parse file here
	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

// Let's guard against something that might panic
func safeValue(vals []int, index int) (n int, err error) { // note  nameed return vals
	defer func() { // we can guard against panics INSIDE a defer anonymous func
		if e := recover(); e != nil {
			err = fmt.Errorf("%v\n", e) // assignment to named return value
		}
	}()

	return vals[index], nil
}

// CHALLENGE -- Kill server PID
func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	// simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)

	if err := os.Remove(pidFile); err != nil {
		// we can go on if we fail here
		log.Printf("warning: can't remove pid file - %s", err)
	}
	return nil
}
