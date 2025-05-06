package models

type RegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RegisterMessage struct {
	Message string      `json:"message"`
	User    *MessageDTO `json:"user"`
}
