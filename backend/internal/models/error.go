package models

type ErrorResponse struct {
	Massage string `json:"massage"`
	Error   error  `json:"error"`
}

type SucessResponse struct {
	Massage string `json:"massage"`
}
