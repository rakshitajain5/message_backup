package models

type ErrorResponse struct {
	Code string `json:"code"`
	Error string `json:"error"`
	Status int `json:"status"`
}

//func (e *ErrorResponse) SetCode(code string) {
//	e.Code = code
//}
//
//func (e *ErrorResponse) SetError(err string) {
//	e.Error = err
//}