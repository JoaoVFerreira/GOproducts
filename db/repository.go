package db

import (
	"encoding/json"
	"log"
	"os"

	"github.com/JoaoVFerreira/GOproducts/httpProduct"
)


func getData() ([]httpProduct.Product, error) {
	var products []httpProduct.Product
	data, err := os.ReadFile("./db/db.json")
	if err != nil {
		log.Fatal("Error reading file:", err)
		return nil, err
	}
	
	err = json.Unmarshal(data, &products) 
	if err != nil {
		log.Fatal("Error parsing data:", err)
		return nil, err
	}
	return products, nil
}

func GetAll() ([]httpProduct.Product, error) {
	products, err := getData(); if err != nil {
		return nil, err
	}

	return products, nil
}

func GetOne(id int) (*httpProduct.Product, error) {
	products, err := getData(); if err != nil {
		return nil, err
	}

	for p := range products {
		if products[p].ID == id {
			return &products[p], nil
		}
	}
	
	return nil, err
}