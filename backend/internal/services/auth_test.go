package services_test

import (
	"errors"
	"testing"

	"back/internal/models"
	"back/internal/services"
	"back/internal/utils"
)

type fakeUserRepo struct {
	users map[string]*models.User
}

func (f *fakeUserRepo) Create(user *models.User) error {
	if user.Email == "error@example.com" {
		return errors.New("create error")
	}
	f.users[user.Email] = user
	return nil
}

func (f *fakeUserRepo) GetByEmail(email string) (*models.User, error) {
	user, ok := f.users[email]
	if !ok {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (f *fakeUserRepo) GetByID(id int) (*models.User, error)          { return nil, nil }
func (f *fakeUserRepo) GetAll() ([]models.User, error)                { return nil, nil }
func (f *fakeUserRepo) UpdateStorageLimit(userID, newLimit int) error { return nil }
func (f *fakeUserRepo) UpdateAdmin(userID int, isAdmin bool) error    { return nil }

func TestRegister(t *testing.T) {
	repo := &fakeUserRepo{users: make(map[string]*models.User)}
	svc := services.NewAuthService(repo)

	dto := models.RegisterDTO{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "pass123",
	}

	user, err := svc.Register(dto)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.Email != dto.Email {
		t.Errorf("expected email %s, got %s", dto.Email, user.Email)
	}

	if user.IsAdmin {
		t.Errorf("new user should not be admin")
	}
}

func TestRegister_FailOnCreate(t *testing.T) {
	repo := &fakeUserRepo{users: make(map[string]*models.User)}
	svc := services.NewAuthService(repo)

	dto := models.RegisterDTO{
		Email:    "error@example.com",
		Name:     "Fail User",
		Password: "pass123",
	}

	_, err := svc.Register(dto)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestLogin_Success(t *testing.T) {
	repo := &fakeUserRepo{users: make(map[string]*models.User)}
	svc := services.NewAuthService(repo)

	hashedPassword, err := utils.HashPassword("password")
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	user := &models.User{
		Email:    "login@example.com",
		Password: hashedPassword,
		IsAdmin:  false,
	}

	err = repo.Create(user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	dto := models.LoginDTO{
		Email:    "login@example.com",
		Password: "password",
	}

	token, u, err := svc.Login(dto)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token == "" {
		t.Error("expected token, got empty string")
	}

	if u.Email != dto.Email {
		t.Errorf("expected email %s, got %s", dto.Email, u.Email)
	}
}

func TestLogin_FailInvalidPassword(t *testing.T) {
	repo := &fakeUserRepo{users: make(map[string]*models.User)}
	svc := services.NewAuthService(repo)

	user := &models.User{
		Email:    "badpass@example.com",
		Password: "$2a$10$7EqJtq98hPqEX7fNZaFWoOQFzNHw4nRXKd6vlY5Lx3cQ7xcl5rDvy",
		IsAdmin:  false,
	}
	repo.Create(user)

	dto := models.LoginDTO{
		Email:    "badpass@example.com",
		Password: "wrongpass",
	}

	_, _, err := svc.Login(dto)
	if err == nil {
		t.Fatal("expected error on bad password, got nil")
	}
}

func TestLogin_FailUserNotFound(t *testing.T) {
	repo := &fakeUserRepo{users: make(map[string]*models.User)}
	svc := services.NewAuthService(repo)

	dto := models.LoginDTO{
		Email:    "missing@example.com",
		Password: "anything",
	}

	_, _, err := svc.Login(dto)
	if err == nil {
		t.Fatal("expected error on missing user, got nil")
	}
}
