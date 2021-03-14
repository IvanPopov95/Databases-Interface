package main

import (
	"fmt"
	"log"
	"net/http"
	"projectttt/controllers"
	"projectttt/psqldb"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func initViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func main() {

	if err := initViper(); err != nil {
		log.Fatal("Cannot init viper", err)
	}

	db, err := psqldb.InitDataBase()
	if err != nil {
		log.Fatal("Error when init database", err)
	}
	handler := controllers.NewHandler(db)
	r := mux.NewRouter()
	r.HandleFunc("/", handler.GetItemsList).Methods("GET")
	r.HandleFunc("/", handler.AddItem).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", handler.GetItemWithID).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", handler.DeleteItem).Methods("DELETE")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)

}
