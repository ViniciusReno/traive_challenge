package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ViniciusReno/traive/internal/controllers"
	"github.com/ViniciusReno/traive/internal/models"
	accountrepo "github.com/ViniciusReno/traive/internal/repository"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)

const (
	limitForErr = 10
	pageForErr  = 1
)

type HttpHandler struct {
	repository *accountrepo.Repository
}

func NewHttpHandler(repository *accountrepo.Repository) *HttpHandler {
	return &HttpHandler{
		repository: repository,
	}
}

func (handler *HttpHandler) CreateTransactions(w http.ResponseWriter, r *http.Request) {
	var req []models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logrus.Warn(err.Error())
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userid := getUserID(r)
	uuid, err := uuid.FromString(userid)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(req); i++ {
		req[i].UserID = uuid
		if err := models.ParseOperationType(string(req[i].OperationType)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := models.ParseOrigin(string(req[i].Origin)); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	tc := controllers.NewTransactionController(handler.repository)
	statuscode, err := tc.CreateTransactions(req)
	if err != nil {
		http.Error(w, err.Error(), statuscode)
		return
	}

	w.WriteHeader(statuscode)
}

func (handler *HttpHandler) ListTransfersByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	if ID == "" {
		logrus.Warn("invalid transaction id")
		http.Error(w, "invalid transaction id", http.StatusBadRequest)
		return
	}

	_, err := uuid.FromString(ID)
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	uc := controllers.NewTransactionController(handler.repository)
	tranfers, statuscode, err := uc.ListTransactionsByID(ID)
	if err != nil {
		http.Error(w, err.Error(), statuscode)
		return
	}

	tranfersJSON, err := json.Marshal(tranfers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(tranfersJSON)
}

func (handler *HttpHandler) ListTransactions(w http.ResponseWriter, r *http.Request) {
	page, limit, filters := getParams(r)

	uc := controllers.NewTransactionController(handler.repository)
	tranfers, statuscode, err := uc.ListTransactions(page, limit, filters)
	if err != nil {
		http.Error(w, err.Error(), statuscode)
		return
	}

	tranfersJSON, err := json.Marshal(tranfers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(tranfersJSON)
}

func getParams(r *http.Request) (int, int, map[string]interface{}) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = pageForErr
		logrus.Info(fmt.Errorf("invalid page, using %v", pageForErr))
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = limitForErr
		logrus.Info(fmt.Errorf("invalid limit, using %v", limitForErr))
	}

	filters := make(map[string]interface{})
	queryValues := r.URL.Query()
	for key := range queryValues {
		if key != "page" && key != "limit" {
			filters[key] = queryValues.Get(key)
		}
	}
	return page, limit, filters
}

func getUserID(r *http.Request) string {
	tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	claims, _ := extractClaims(tokenString)
	userID := claims["userid"].(string)

	return userID
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "secret"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
