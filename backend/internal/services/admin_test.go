package services_test

import (
	"errors"
	"testing"

	"back/internal/models"
	"back/internal/services"
)

type fakeUserRepoAdmin struct {
	users        []models.User
	updateCalls  []string
	updateErrors map[string]error
	GetAllFunc   func() ([]models.User, error)
}

func (f *fakeUserRepoAdmin) GetAllAdmins() ([]models.User, error) {
	panic("unimplemented")
}

func (f *fakeUserRepoAdmin) Create(user *models.User) error {
	panic("unimplemented")
}

func (f *fakeUserRepoAdmin) GetByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

func (f *fakeUserRepoAdmin) GetByID(id int) (*models.User, error) {
	panic("unimplemented")
}

func (f *fakeUserRepoAdmin) GetAll() ([]models.User, error) {
	if f.GetAllFunc != nil {
		return f.GetAllFunc()
	}
	return f.users, nil
}

func (f *fakeUserRepoAdmin) UpdateStorageLimit(userID, newLimit int) error {
	if err, ok := f.updateErrors["storage"]; ok {
		return err
	}
	f.updateCalls = append(f.updateCalls, "storage")
	return nil
}

func (f *fakeUserRepoAdmin) UpdateAdmin(userID int, isAdmin bool) error {
	if err, ok := f.updateErrors["admin"]; ok {
		return err
	}
	f.updateCalls = append(f.updateCalls, "admin")
	return nil
}

func TestGetAllUsers(t *testing.T) {
	repo := &fakeUserRepoAdmin{
		users: []models.User{
			{ID: 1, Email: "u1@example.com"},
			{ID: 2, Email: "u2@example.com"},
		},
	}
	svc := services.NewAdminService(repo)

	users, err := svc.GetAllUsers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}

	if users[0].Email != "u1@example.com" {
		t.Errorf("expected first user email u1@example.com, got %s", users[0].Email)
	}
}

func TestUpdateStorageLimit_Valid(t *testing.T) {
	repo := &fakeUserRepoAdmin{}
	svc := services.NewAdminService(repo)

	err := svc.UpdateStorageLimit(1, 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repo.updateCalls) != 1 || repo.updateCalls[0] != "storage" {
		t.Errorf("expected updateCalls to contain 'storage', got %v", repo.updateCalls)
	}
}

func TestUpdateStorageLimit_Invalid(t *testing.T) {
	repo := &fakeUserRepoAdmin{}
	svc := services.NewAdminService(repo)

	err := svc.UpdateStorageLimit(1, 0)
	if err == nil {
		t.Fatal("expected error for invalid storage limit, got nil")
	}
}

func TestMakeAdmin(t *testing.T) {
	repo := &fakeUserRepoAdmin{}
	svc := services.NewAdminService(repo)

	err := svc.MakeAdmin(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repo.updateCalls) != 1 || repo.updateCalls[0] != "admin" {
		t.Errorf("expected updateCalls to contain 'admin', got %v", repo.updateCalls)
	}
}

func TestRevokeAdmin(t *testing.T) {
	repo := &fakeUserRepoAdmin{}
	svc := services.NewAdminService(repo)

	err := svc.RevokeAdmin(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repo.updateCalls) != 1 || repo.updateCalls[0] != "admin" {
		t.Errorf("expected updateCalls to contain 'admin', got %v", repo.updateCalls)
	}
}

func TestGetAllUsers_Error(t *testing.T) {
	repo := &fakeUserRepoAdmin{
		GetAllFunc: func() ([]models.User, error) {
			return nil, errors.New("forced error")
		},
	}

	svc := services.NewAdminService(repo)

	_, err := svc.GetAllUsers()
	if err == nil {
		t.Fatal("expected error from GetAllUsers, got nil")
	}
}

func TestUpdateStorageLimit_ErrorFromRepo(t *testing.T) {
	repo := &fakeUserRepoAdmin{
		updateErrors: map[string]error{
			"storage": errors.New("update error"),
		},
	}
	svc := services.NewAdminService(repo)

	err := svc.UpdateStorageLimit(1, 50)
	if err == nil {
		t.Fatal("expected error from UpdateStorageLimit, got nil")
	}
}

func TestMakeAdmin_ErrorFromRepo(t *testing.T) {
	repo := &fakeUserRepoAdmin{
		updateErrors: map[string]error{
			"admin": errors.New("update admin error"),
		},
	}
	svc := services.NewAdminService(repo)

	err := svc.MakeAdmin(1)
	if err == nil {
		t.Fatal("expected error from MakeAdmin, got nil")
	}
}

func TestRevokeAdmin_ErrorFromRepo(t *testing.T) {
	repo := &fakeUserRepoAdmin{
		updateErrors: map[string]error{
			"admin": errors.New("update admin error"),
		},
	}
	svc := services.NewAdminService(repo)

	err := svc.RevokeAdmin(1)
	if err == nil {
		t.Fatal("expected error from RevokeAdmin, got nil")
	}
}
