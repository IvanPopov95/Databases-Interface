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

// NewHandler return handler for handlefunc
func NewHandler(db *sql.DB) Handler {
	return Handler{DB: db}
}

// GetItemsListController return all items
func (h Handler) GetItemsListController(w http.ResponseWriter, req *http.Request) {
	items, err := psqldb.GetItemsList(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(itemsJSON)

}

// GetItemWithIDController get one item with id from path params
func (h Handler) GetItemWithIDController(w http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)["id"]
	id, err := strconv.Atoi(pathParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	item, err := psqldb.GetItemWithID(h.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	itemJSON, err := json.Marshal(*item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(itemJSON)
}

// AddItemController adding item
func (h Handler) AddItemController(w http.ResponseWriter, req *http.Request) {
	// name := req.URL.Query()["name"][0]
	var m models.Item
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

	err = psqldb.AddItem(h.DB, m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteItemController delete item with id
func (h Handler) DeleteItemController(w http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)["id"]
	id, err := strconv.Atoi(pathParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = psqldb.DeleteItem(h.DB, id)
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
