package postgres

import (
	"database/sql"

	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/model"
	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/repository"
)

type PGUserRepo struct {
	db *sql.DB
}

func NewPGUserRepo(db *sql.DB) repository.UserRepository {
	return PGUserRepo{db}
}

func (r *PGUserRepo) GetByID(id int64) (*model.User, error) {
	user := model.User{}
	err := r.db.QueryRow(
		"SELECT ID, NAME, EMAIL FROM users WHERE ID = $1", id,
	).Scan(user.ID, user.Name, user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PGUserRepo) Create(user *model.User) error {
	err := r.db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		user.Name, user.Email,
	).Scan(user.ID)

	if err != nil {
		return err
	}
	return nil
}
