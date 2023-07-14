package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// ...for quick debugging
var pl = fmt.Println

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err)
	err = parsedTemplate.Execute(w, nil)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func getSum(x, y int) int {
	return x + y
}

func getQuotient(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot devide by zero")
		return 0, err
	} else {
		return (x / y), nil
	}
}

func addHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello internet gophers")
	sum := getSum(5, 4)
	output := fmt.Sprintf("5 + 4 = %d\n", sum)
	write(writer, output)
}
func divideHandler(writer http.ResponseWriter, request *http.Request) {
	v, err := getQuotient(5, 0)
	if err != nil {
		write(writer, "can't divide by zero\n")
	} else {
		output := fmt.Sprintf("5 / 4 = %.2f\n", v)
		write(writer, output)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getsum", addHandler)
	http.HandleFunc("/divide", divideHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}
