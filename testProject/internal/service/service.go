package service

import (
	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/model"
	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUser(id int64) (*model.User, error) {
	return s.repo.GetByID(id)
}
