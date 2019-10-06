package main

import (
	"log"

	"github.com/rickihastings/shopping-cart/pkg/api"
)

func main() {
	log.Println("Starting up shopping-cart API...")

	api.CreateHTTPServer()
}
