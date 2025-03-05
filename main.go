package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/split", handleSplit)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func handleSplit(w http.ResponseWriter, r *http.Request) {
	// Parse form data, process EPUB/text, and render results
}