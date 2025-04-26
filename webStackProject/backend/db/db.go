
package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "os"
)

func Connect() (*sql.DB, error) {
    connStr := os.Getenv("POSTGRES_CONN")
    return sql.Open("postgres", connStr)
}
