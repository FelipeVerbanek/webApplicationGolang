package db

import "database/sql"

func ConnectDB() *sql.DB {
	connect := "user=postgres dbname=loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err.Error())
	}

	return db
}
