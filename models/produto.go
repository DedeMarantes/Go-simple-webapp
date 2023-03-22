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
	allProducts, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
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
		produto.Id = id
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

func DeletaProduto(id string) {
	db := db.ConnectDb()
	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConnectDb()
	produto, err := db.Query("SELECT * from produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	novoProduto := Produto{}

	for produto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err := produto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		novoProduto.Id = id
		novoProduto.Nome = nome
		novoProduto.Descricao = descricao
		novoProduto.Preco = preco
		novoProduto.Quantidade = quantidade
	}
	defer db.Close()
	return novoProduto
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()
	atualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
