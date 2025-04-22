package model

import "database/sql"

type User struct {
	ID    int64          `json:"id"`
	Name  string         `json:"name"`
	Email sql.NullString `json:"email"`
}
