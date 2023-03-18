package main

import (
	"net/http"
	"text/template"
)

var templ = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index", nil)
}
