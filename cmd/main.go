package main

import (
	"log"

	"github.com/vivekup3424/shopping-cart/cmd/api"
)

func main() {
	const addr = "localhost:8080"
	server := api.NewAPIServer(addr, nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
