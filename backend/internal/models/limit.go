package models

type UpdateStorage struct {
	UserID   int `json:"user_id"`
	NewLimit int `json:"new_limit"`
}
