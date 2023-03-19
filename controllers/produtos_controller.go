package controllers

import (
	"html/template"
	"log"
	"modulo/models"
	"net/http"
	"strconv"
)

var templ = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.GetAllProducts()
	templ.ExecuteTemplate(w, "index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convercao do preco: ", err)
		}
		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na convercao da quantidade: ", err)
		}

		models.CreateProduct(nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", 301)
}
