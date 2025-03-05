package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/split", handleSplit)
	http.ListenAndServe(":"+port, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func handleSplit(w http.ResponseWriter, r *http.Request) {
	// Parse form data, process EPUB/text, and render results
}