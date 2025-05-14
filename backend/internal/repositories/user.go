package repositories

import (
	"back/internal/models"

	log "github.com/charmbracelet/log"

	"github.com/jmoiron/sqlx"
)

const (
	CreateUserSql         = `INSERT INTO users (email, name, password, is_admin, storage_limit) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	GetUserByEmailSql     = `SELECT * FROM users WHERE email = $1`
	GetUserByIdSql        = `SELECT * FROM users WHERE id = $1`
	GetAllUsersSql        = `SELECT id, email, name, is_admin, storage_limit FROM users`
	UpdateStorageLimitSql = `UPDATE users SET storage_limit = $1 WHERE id = $2`
	UpdateAdminSql        = `UPDATE users SET is_admin = $1 WHERE id = $2`
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(user *models.User) error {
	return r.DB.QueryRow(
		CreateUserSql,
		user.Email,
		user.Name,
		user.Password,
		user.IsAdmin,
		user.StorageLimit,
	).Scan(&user.ID)
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Get(&user, GetUserByEmailSql, email)
	if err != nil {
		log.Error("get by email error", "err", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.DB.Get(&user, GetUserByIdSql, id)
	if err != nil {
		log.Error("get by id user", "err", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Select(&users, GetAllUsersSql)
	if err != nil {
		log.Error("get all", "err", err)
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) UpdateStorageLimit(userID, newLimit int) error {
	result, err := r.DB.Exec(UpdateStorageLimitSql, newLimit, userID)
	log.Debug("database update storage limit", result)
	if err != nil {
		log.Error("update storage limit", "err", err)
		return err
	}
	return nil
}

func (r *UserRepo) UpdateAdmin(userID int, isAdmin bool) error {
	result, err := r.DB.Exec(UpdateAdminSql, isAdmin, userID)
	log.Debug("database update admin", result)
	if err != nil {
		log.Error("update update admin", "err", err)
		return err
	}
	return nil
}
