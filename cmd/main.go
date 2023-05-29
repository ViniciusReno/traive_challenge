package main

import (
	"log"
	"net/http"

	db "github.com/ViniciusReno/traive/internal/db"
	httpRouter "github.com/ViniciusReno/traive/internal/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}

func main() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	r.Use(c.Handler)

	db, err := db.ConnectDB()
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()

	httpRouter.Register(r, db)

	log.Print("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Errorf("Error starting server: %v", err)
	}
}
