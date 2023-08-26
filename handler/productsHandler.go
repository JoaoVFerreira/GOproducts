package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JoaoVFerreira/GOproducts/db"
	"github.com/JoaoVFerreira/GOproducts/httpProduct"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetAll(); if err != nil {
		http.Error(w, "Data not found", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")

	response := httpProduct.Response{
		Message: "Product find with success!",
		StatusCode: http.StatusOK,
		Response: products,
	}

	responseJson, err := json.Marshal(&response); if err != nil {
		http.Error(w, "Error parsing data", http.StatusInternalServerError)
	}

	w.Write(responseJson)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(idParam); if err != nil {
		http.Error(w, "Param is not processable", http.StatusUnprocessableEntity)
	}

	product, err := db.GetOne(id); if err != nil {
		http.Error(w, "Param is not processable", http.StatusNotFound)
	}

	response := httpProduct.Response{
		Message: "Product find with success!",
		StatusCode: http.StatusOK,
		Response: product,
	}

	responseJson, err := json.Marshal(&response); if err != nil {
		http.Error(w, "Error parsing data", http.StatusInternalServerError)
	}

	w.Write(responseJson)
}