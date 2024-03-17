package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Divyajothi")
	})
	http.HandleFunc("/Rollno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212CT109")
	})
	http.HandleFunc("/Course", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Computer Technology")
	})
	http.HandleFunc("/year", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "III")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/hello\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}