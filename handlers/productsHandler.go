package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/JoaoVFerreira/GOproducts/db"
	httpProduct "github.com/JoaoVFerreira/GOproducts/http"
)

const errorParsingData = "Error parsing data"
const errorCreatingData = "Error when trying to creating data"
const notProcessableParam = "Param is not processable"


func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetAll(); if err != nil {
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := httpProduct.Response{
		Message: "Product find with success!",
		StatusCode: http.StatusOK,
		Response: products,
	}

	responseJson, err := json.Marshal(&response); if err != nil {
		http.Error(w, errorParsingData, http.StatusInternalServerError)
		return
	}

	w.Write(responseJson)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(idParam); if err != nil {
		http.Error(w, notProcessableParam, http.StatusUnprocessableEntity)
		return
	}

	product, err := db.GetOne(id); if err != nil {
		http.Error(w, notProcessableParam, http.StatusNotFound)
		return
	}

	response := httpProduct.Response{
		Message: "Product find with success!",
		StatusCode: http.StatusOK,
		Response: product,
	}

	responseJson, err := json.Marshal(&response); if err != nil {
		http.Error(w, errorParsingData, http.StatusInternalServerError)
		return
	}

	w.Write(responseJson)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body); if err != nil {
		http.Error(w, "Body is required", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var p httpProduct.Product

	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, errorParsingData, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	p.Validate(w)

	product, err := db.Create(p); if err != nil {
		http.Error(w, errorCreatingData, http.StatusInternalServerError)
		return
	}

	response := httpProduct.Response{
		Message: "Product created with success!",
		StatusCode: http.StatusOK,
		Response: product,
	}

	responseJson, err := json.Marshal(&response); if err != nil {
		http.Error(w, errorParsingData, http.StatusInternalServerError)
	}
	w.Write(responseJson)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(idParam); if err != nil {
		http.Error(w, notProcessableParam, http.StatusUnprocessableEntity)
		return
	}

	idRemoved, err := db.Delete(id); if err != nil {
		http.Error(w, notProcessableParam, http.StatusNotFound)
		return
	}

	response := httpProduct.Response{
		Message: "Product deleted",
		StatusCode: http.StatusOK,
		Response: idRemoved,
	}

	responseJson, err := json.Marshal(&response); if err != nil {
		http.Error(w, errorParsingData, http.StatusInternalServerError)
	}

	w.Write(responseJson)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")
	body, err := io.ReadAll(r.Body); if err != nil {
		http.Error(w, "Body is required", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var p httpProduct.Product

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(idParam); if err != nil {
		http.Error(w, notProcessableParam, http.StatusUnprocessableEntity)
		return
	}

	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, errorParsingData, http.StatusInternalServerError)
		return
	}

	product, err := db.Update(id, p); if err != nil {
		http.Error(w, notProcessableParam, http.StatusNotFound)
		return
	}

	response := httpProduct.Response{
		Message: "Product updated with success!",
		StatusCode: http.StatusOK,
		Response: product,
	}

	responseJson, err := json.Marshal(&response); if err != nil {
		http.Error(w, errorParsingData, http.StatusInternalServerError)
		return
	}

	w.Write(responseJson)
}