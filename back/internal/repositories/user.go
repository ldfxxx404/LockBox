package repositories

import (
	"back/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(user *models.User) error {
	query := `INSERT INTO users (email, name, password, is_admin, storage_limit) VALUES (:email, :name, :password, :is_admin, :storage_limit)`
	_, err := r.DB.NamedExec(query, user)
	return err
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = ?`
	err := r.DB.Get(&user, query, email)
	return &user, err
}

func (r *UserRepo) GetByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = ?`
	err := r.DB.Get(&user, query, id)
	return &user, err
}

func (r *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User
	query := `SELECT id, email, name, is_admin FROM users`
	err := r.DB.Select(&users, query)
	return users, err
}

func (r *UserRepo) UpdateStorageLimit(userID, newLimit int) error {
	query := `UPDATE users SET storage_limit = ? WHERE id = ?`
	_, err := r.DB.Exec(query, newLimit, userID)
	return err
}

func (r *UserRepo) UpdateAdmin(userID int, isAdmin bool) error {
	query := `UPDATE users SET is_admin = ? WHERE id = ?`
	_, err := r.DB.Exec(query, isAdmin, userID)
	return err
}
