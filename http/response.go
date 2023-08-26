package http

type Response struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Response   interface{} `json:"response"`
}