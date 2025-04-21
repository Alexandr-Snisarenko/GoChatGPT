package main

import (
	"fmt"

	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/config"
	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/db"
	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/repository/postgres"
	"github.com/Alexandr-Snisarenko/GoChatGPT/testProgect/internal/service"
)

func main() {
	cfg := config.NewConfig()
	database := db.NewPostgres(cfg)
	defer database.Close()

	rep := postgres.NewPGUserRepo(database)

	userServ := service.NewUserService(rep)

	if err := userServ.Start(); err != nil {
		fmt.Printf("Server end with error %s", err)
	} else {
		fmt.Println("Service done")
	}
}
