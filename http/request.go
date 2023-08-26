package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type rating struct {
	Rate  float32 `json:"rate"`
	Count int     `json:"count"`
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Rating      *rating `json:"rating"`
}

func checkProductParams(param interface{}, typ string, w http.ResponseWriter) {
	r := Response{
		Message: "Wrong use of sent body params",
		StatusCode: http.StatusUnprocessableEntity,
		Response: fmt.Sprintf("Parameter %s (type %s) could not be empty", param, typ),
	}

	response, err := json.Marshal(&r); if err != nil {
		http.Error(w, "Error parsing data", http.StatusInternalServerError)
	}

	w.Write(response)
}

func (p *Product) Validate(w http.ResponseWriter) {
	if p.Title == "" {
		checkProductParams(p.Title, "string", w)
	}
	if p.Price == 0 {
		checkProductParams(p.Price, "float32", w)
	}
	if p.Description == "" {
		checkProductParams(p.Description, "string", w)
	}
	if p.Category == "" {
		checkProductParams(p.Category, "string", w)
	}
	if p.Rating.Count == 0 {
		checkProductParams(p.Rating.Count, "int", w)
	}
	if p.Rating.Rate == 0 {
		checkProductParams(p.Rating.Rate, "float32", w)
	}
}