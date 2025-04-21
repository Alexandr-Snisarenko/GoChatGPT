package service

import (
	"database/sql"

	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/model"
	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Start() error {
	user := model.User{ID: 0, Name: "Test", Email: sql.NullString{Valid: false}}
	err := s.repo.Create(&user)
	return err
}
