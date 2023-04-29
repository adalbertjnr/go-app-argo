package controllers

import (
	"fmt"
	"go-app/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.SelectProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Problema ao converter preco para float:", err)
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro ao converter quantidade para int:", err)
		}
		models.CriarProduto(nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeleteId(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produtoEditado := models.EditId(idProduto)
	temp.ExecuteTemplate(w, "Edit", produtoEditado)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Erro ao converter stringId para Int:", err)
		}

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Erro ao converter stringPreco para float64:", err)
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro ao converter stringQuantidade para inteiro:", err)
		}

		models.UpdateTable(idInt, nome, descricao, precoFloat, quantidadeInt)

	}

	http.Redirect(w, r, "/", 301)
}
