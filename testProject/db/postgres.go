package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/config"
	_ "github.com/jackc/pgx/stdlib"
)

func NewPostgres(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("База недоступна:", err)
	}

	return db
}
