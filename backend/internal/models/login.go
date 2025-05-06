package models

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginMessage struct {
	Message string      `json:"message"`
	User    *MessageDTO `json:"user"`
	Token   string      `json:"token"`
}
