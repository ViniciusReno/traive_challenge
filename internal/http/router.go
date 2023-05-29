package http

import (
	"net/http"

	"github.com/ViniciusReno/traive/internal/auth"
	accountrepo "github.com/ViniciusReno/traive/internal/repository"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Register(router *mux.Router, database *gorm.DB) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	repository := accountrepo.NewRepository(database)
	handler := NewHttpHandler(repository)

	router.HandleFunc("/account/person", handler.CreatePersonHandler).Methods("POST")
	router.HandleFunc("/account/login", handler.LoginHandler).Methods("POST")

	authenticated := router.PathPrefix("").Subrouter()
	authenticated.Use(auth.JWTMiddleware)
	authenticated.HandleFunc("/account/transaction", handler.ListTransactions).Methods("GET")
	authenticated.HandleFunc("/account/transaction/{id}", handler.ListTransfersByID).Methods("GET")
	authenticated.HandleFunc("/account/transaction", handler.CreateTransactions).Methods("POST")
}
