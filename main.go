package main

import (
	"log"
	"os"
	"posts/internals/handler"
	"posts/internals/repository"
	"posts/internals/server"
	"posts/internals/service"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to init DB: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	serv := service.NewService(rep)
	handl := handler.NewHandler(serv)

	srv := new(server.Server)
	log.Fatal(srv.Run(":8080", handl.InitRouter()))
}
