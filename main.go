package main

import (
	"fmt"
	"net/http"

	"github.com/JoaoVFerreira/GOproducts/handler"
)

func main() {
	// Mux work as a router
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/products", handler.GetProducts)
	mux.HandleFunc("/api/v1/product", handler.GetOneProduct)

	// Server listening on 3000 
	fmt.Println("Server listening on 3000")
	http.ListenAndServe(":3000", mux)
}

