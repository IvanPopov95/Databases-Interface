package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler struct
type Handler struct {
	DB *sql.DB
}

// Item main struct
type Item struct {
	ID   int
	Name string
}

// NewHandler return handler for handlefunc
func NewHandler(db *sql.DB) Handler {
	return Handler{DB: db}
}

// GetItemsList return all items
func (h Handler) GetItemsList(w http.ResponseWriter, req *http.Request) {
	rows, err := h.DB.Query("SELECT id, name FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}
	rows.Close()
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(itemsJSON)

}

// GetItemWithID get one item with id from path params
func (h Handler) GetItemWithID(w http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)["id"]
	id, err := strconv.Atoi(pathParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var item Item
	err = h.DB.QueryRow("SELECT * FROM items WHERE id = $1", id).Scan(&item.Name, &item.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	itemJSON, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(itemJSON)
}

// AddItem adding item
func (h Handler) AddItem(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query()["name"][0]
	_, err := h.DB.Exec("INSERT INTO items(name) values($1)", name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteItem delete item with id
func (h Handler) DeleteItem(w http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)["id"]
	id, err := strconv.Atoi(pathParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = h.DB.Exec("delete from items where id=$1", id)
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
