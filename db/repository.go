package db

import (
	"encoding/json"
	"log"
	"os"

	"github.com/JoaoVFerreira/GOproducts/request"
)


func getData() ([]request.Product, error) {
	var products []request.Product
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

func GetAll() ([]request.Product, error) {
	products, err := getData(); if err != nil {
		return nil, err
	}

	return products, nil
}

func GetOne(id int) (*request.Product, error) {
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