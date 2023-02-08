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

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{"Camiseta", "Confortavel", 38, 3},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Confortavel", 59, 3},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
