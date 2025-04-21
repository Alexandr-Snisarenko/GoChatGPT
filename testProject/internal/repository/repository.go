package repository

import "github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/model"

type UserRepository interface {
	GetByID(id int64) (*model.User, error)
	Create(user *model.User) error
}
