package models

import (
	"go-app/db"
	"log"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SelectProdutos() []Produto {
	db := db.ConectaBanco()

	selectProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()
	insertBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1,$2,$3,$4)")
	if err != nil {
		log.Fatal(err)
	}
	insertBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteId(idProduto string) {
	db := db.ConectaBanco()
	deleteFromBanco, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		log.Fatal(err)
	}
	deleteFromBanco.Exec(idProduto)
	defer db.Close()
}

func EditId(idProduto string) Produto {
	db := db.ConectaBanco()
	selectProduto, err := db.Query("select * from produtos where id=$1", idProduto)
	if err != nil {
		log.Fatal(err)
	}

	produtoAtualizado := Produto{}

	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			log.Fatal(err)
		}
		produtoAtualizado.Id = id
		produtoAtualizado.Nome = nome
		produtoAtualizado.Descricao = descricao
		produtoAtualizado.Preco = preco
		produtoAtualizado.Quantidade = quantidade

	}
	defer db.Close()
	return produtoAtualizado
}

func UpdateTable(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()
	updateToDatabase, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		log.Fatal(err)
	}

	updateToDatabase.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
