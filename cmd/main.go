package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/vivekup3424/shopping-cart/cmd/api"
	"github.com/vivekup3424/shopping-cart/db"
)

func main() {
	const addr = "localhost:8080"
	//I can move this detail to a environment variable
	db, _ := db.MySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "localhost:2432",
		DBName:               "shopping_cart",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(addr, nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
