package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func connectDb() *sql.DB {
	conexao := "user=root dbname=postgres password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templ = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	db := connectDb()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDb()
	allProducts, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	produto := Produto{}
	lista_produtos := []Produto{}
	//escanera banco de dados
	for allProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = allProducts.Scan(&id, &nome, &descricao, &preco, &quantidade) //pegar informações
		if err != nil {
			panic(err.Error())
		}
		//adicionar produtos
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		lista_produtos = append(lista_produtos, produto)
	}

	templ.ExecuteTemplate(w, "index", lista_produtos)
	defer db.Close()
}
