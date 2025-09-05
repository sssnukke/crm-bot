package main

import (
	"log"
	"net/http"
	_ "os"

	"back/internal/config"
	"back/internal/db"
	"back/internal/transport/http2"
)

func main() {
	cfg := config.Load()

	database := db.Connect(cfg.DatabaseURL)

	router := http2.NewRouter(cfg, database)

	log.Printf("Server started on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
