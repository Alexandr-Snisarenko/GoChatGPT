package mock

import (
	"database/sql"

	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/model"
	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/repository"
)

type MockUserRepo struct{}

func NewMockUserRepo() repository.UserRepository {
	return &MockUserRepo{}
}

func (m *MockUserRepo) GetByID(id int64) (*model.User, error) {
	return &model.User{ID: id, Name: "Test User", Email: sql.NullString{String: "test@example.com", Valid: true}}, nil
}

func (m *MockUserRepo) Create(user *model.User) error {
	return nil
}
