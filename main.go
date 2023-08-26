package main

import (
	"fmt"
	"net/http"

	"github.com/JoaoVFerreira/GOproducts/handlers"
)

func main() {
	// Mux work as a router
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/products", handlers.GetProducts)
	mux.HandleFunc("/api/v1/product", handlers.GetOneProduct)
	mux.HandleFunc("/api/v1/products/create", handlers.CreateProduct)

	// Server listening on 3000 
	fmt.Println("Server listening on 3000")
	http.ListenAndServe(":3000", mux)
}

