package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templ = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, manga curta", Preco: 39, Quantidade: 4},
		{Nome: "Fone", Descricao: "Fone de celular", Preco: 45, Quantidade: 10},
		{Nome: "Pendrive", Descricao: "pendrive 32gb", Preco: 20, Quantidade: 4},
	}
	templ.ExecuteTemplate(w, "index", produtos)
}
