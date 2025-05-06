package models

type MessageDTO struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type User struct {
	ID           int    `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	Name         string `db:"name" json:"name"`
	Password     string `db:"password" json:"-"`
	IsAdmin      bool   `db:"is_admin" json:"is_admin"`
	StorageLimit int    `db:"storage_limit" json:"storage_limit"`
}
