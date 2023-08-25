package db

import (
	"encoding/json"
	"log"
	"os"

	"github.com/JoaoVFerreira/GOproducts/handler"
)

func GetAllData() ([]handler.Product, error) {
	var products []handler.Product
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