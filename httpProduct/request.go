package httpProduct

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