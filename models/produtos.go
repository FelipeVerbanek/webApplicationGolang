package models

import "github.com/felipeverbanek/app-web/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindAllProducts() []Produto {
	db := db.ConnectDB()
	selectAllProducts, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
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

func CreateProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()

	insertData, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProductForId, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProductForId.Exec(id)
	defer db.Close()
}

func FindProduct(id string) Produto {
	db := db.ConnectDB()

	selectProduct, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Produto{}
	for selectProduct.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProduct.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Nome = nome
		product.Descricao = descricao
		product.Preco = preco
		product.Quantidade = quantidade
	}

	defer db.Close()
	return product
}

func UpdateProduct(id int, nome, descricao string, quantidade int, preco float64) {
	db := db.ConnectDB()

	updateData, err := db.Prepare("update produtos set nome = $1, descricao = $2, quantidade = $3, preco = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	updateData.Exec(nome, descricao, quantidade, preco, id)
	defer db.Close()
}
