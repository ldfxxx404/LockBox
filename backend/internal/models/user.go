package models

type User struct {
	ID           int    `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	Name         string `db:"name" json:"name"`
	Password     string `db:"password" json:"-"`
	IsAdmin      bool   `db:"is_admin" json:"is_admin"`
	StorageLimit int    `db:"storage_limit" json:"storage_limit"`
}

type UserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MassageDTO struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type RegisterMassage struct {
	Message string      `json:"message"`
	User    *MassageDTO `json:"user"`
}

type LoginMassage struct {
	Message string      `json:"message"`
	User    *MassageDTO `json:"user"`
	Token   string      `json:"token"`
}
