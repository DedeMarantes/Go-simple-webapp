package models

import (
	"modulo/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Produto {
	db := db.ConnectDb()
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
	defer db.Close()
	return lista_produtos
}

func CreateProduct(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()
	insertData, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
