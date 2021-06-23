package main

import (
	"database/sql"
	db2 "github.com/turnes/hexagonal-architecture/adapters/db"
	"github.com/turnes/hexagonal-architecture/app"
)

func main() {
	var db, _ = sql.Open("sqlite3","db.sqlite")
	productDb := db2.NewProductDb(db)
	service := app.NewProductService(productDb)
	service.Create("Product1", 30.5)

}