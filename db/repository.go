package db

import (
	"encoding/json"
	"log"
	"os"

	httpProduct "github.com/JoaoVFerreira/GOproducts/http"
)

const errorParsingData = "Error parsing data:"
const dbPath = "./db/db.json"

func getData() ([]httpProduct.Product, error) {
	var products []httpProduct.Product
	data, err := os.ReadFile(dbPath); if err != nil {
		log.Fatal("Error reading file:", err)
		return nil, err
	}
	
	err = json.Unmarshal(data, &products); if err != nil {
		log.Fatal(errorParsingData, err)
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

func Create(p httpProduct.Product) (* httpProduct.Product, error) {
	products, err := getData(); if err != nil {
		return nil, err
	}

	p.ID = len(products) + 1
	products = append(products, p)
	pJson, err := json.Marshal(&products); if err != nil {
		log.Fatal(errorParsingData, err)
	}

	os.WriteFile(dbPath, pJson, 0666)
	return &p, nil
}

func Delete(id int) (int, error) {
	products, err := getData(); if err != nil {
		return 0, err
	}

	idToRemove := -1
	for i, p := range products {
		if p.ID == id {
			idToRemove = i
			break
		}
	}

	if idToRemove != -1 {
		products = append(products[:idToRemove], products[idToRemove+1:]...)
		pJson, err := json.Marshal(&products); if err != nil {
			log.Fatal(errorParsingData, err)
		}
		os.WriteFile(dbPath, pJson, 0666)
		return id, nil
	} 
	return 0, nil
}

func Update(id int, p httpProduct.Product) (* httpProduct.Product, error) {
	products, err := getData()
	if err != nil {
		return nil, err
	}

	var foundProduct *httpProduct.Product

	for i := range products {
		if products[i].ID == id {
			foundProduct = &products[i]
			break
		}
	}
	if p.Title != "" {
		foundProduct.Title = p.Title
	}
	if p.Price != 0 {
		foundProduct.Price = p.Price
	}
	if p.Description != "" {
		foundProduct.Description = p.Description
	}
	if p.Category != "" {
		foundProduct.Category = p.Category
	}
	if p.Rating != nil {
		foundProduct.Rating = p.Rating
	}
	if foundProduct == nil {
		return nil, err
	}
	
	updatedData, err := json.MarshalIndent(products, "", "    ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(dbPath, updatedData, 0644)
	if err != nil {
		return nil, err
	}

	return foundProduct, nil
}