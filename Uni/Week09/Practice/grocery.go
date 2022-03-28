package main

import (
	"log"
	"net/http"
	"html/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeList(writer http.ResponseWriter, list []string) {
	html_inner, err := template.ParseFiles("list.html")
	check(err)
	// execute
	err = html_inner.Execute(writer, list)
	// after execute the check
	check(err)
}

func fruitHandler(writer http.ResponseWriter, request *http.Request) {
	writeList(writer, []string{"apples", "oranges", "pears"})
}

func meatHandler(writer http.ResponseWriter, request *http.Request) {
	writeList(writer, []string{"chicken", "beef", "lamb"})
}

func main() {
	http.HandleFunc("/fruit", fruitHandler)
	http.HandleFunc("/meat", meatHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
