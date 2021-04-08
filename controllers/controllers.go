package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projectttt/models"
	"projectttt/psqldb"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler struct
type Handler struct {
	DB *sql.DB
}

// Database - interface for different databases
type Database interface {
	GetUsersList() ([]models.User, error)
	GetUserWithID(id int) (*models.User, error)
	AddUser(models.User) error
	DeleteUser(id int) error
}

// NewHandler return handler for handlefunc
func NewHandler(db *sql.DB) Handler {
	return Handler{DB: db}
}

// GetUsersListController return all users
func (h Handler) GetUsersListController(w http.ResponseWriter, req *http.Request) {
	users, err := psqldb.GetUsersList(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(usersJSON)

}

// GetUserWithIDController get one user with id from path params
func (h Handler) GetUserWithIDController(w http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)["id"]
	id, err := strconv.Atoi(pathParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := psqldb.GetUserWithID(h.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userJSON, err := json.Marshal(*user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(userJSON)
}

// AddUserController adding user
func (h Handler) AddUserController(w http.ResponseWriter, req *http.Request) {
	var m models.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = psqldb.AddUser(h.DB, m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteUserController delete user with id
func (h Handler) DeleteUserController(w http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)["id"]
	id, err := strconv.Atoi(pathParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = psqldb.DeleteUser(h.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("User %s deleted", pathParams)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseJSON)
}
