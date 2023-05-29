package http

import (
	"encoding/json"
	"net/http"

	"github.com/ViniciusReno/traive/internal/controllers"
	"github.com/ViniciusReno/traive/internal/models"
	"github.com/sirupsen/logrus"
)

func (handler *HttpHandler) CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logrus.Warn(err.Error())
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	uc := controllers.NewUserController(handler.repository)

	token, err := uc.CreatePerson(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(token)
}

func (handler *HttpHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logrus.Warn(err.Error())
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	uc := controllers.NewUserController(handler.repository)
	token, statuscode, err := uc.Login(req)
	if err != nil {
		http.Error(w, err.Error(), statuscode)
		return
	}

	w.WriteHeader(statuscode)
	w.Write(token)
}
