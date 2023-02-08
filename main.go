package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	connect := "user=postgres dbname=loja postgres=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err.Error())
	}

	return db
}

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

	db := connectDB()
	defer db.Close()

	produtos := []Produto{
		{"Camiseta", "Confortavel", 38, 3},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Confortavel", 59, 3},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
