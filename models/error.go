package models

type ErrorResponse struct {
	Code   string `json:"code"`
	Error  string `json:"error"`
	Status int `json:"status"`
}




