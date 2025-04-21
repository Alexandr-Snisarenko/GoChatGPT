package model

import "database/sql"

type User struct {
	ID    int64
	Name  string
	Email sql.NullString
}
