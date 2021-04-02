package psqldb

import (
	"database/sql"
	"fmt"
	"projectttt/models"

	// driver for postgres, use in Open
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
)

// InitDataBase connect to psqldb
func InitDataBase() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("storage.postgres.username"),
		viper.GetString("storage.postgres.password"),
		viper.GetString("storage.postgres.host"),
		viper.GetString("storage.postgres.port"),
		viper.GetString("storage.postgres.dbname"))
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetItemsList get all items from db
func GetItemsList(db *sql.DB) ([]models.Item, error) {
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// GetItemWithID get one item with id
func GetItemWithID(db *sql.DB, id int) (*models.Item, error) {
	var item models.Item
	err := db.QueryRow("SELECT * FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// AddItem adding one item to db
func AddItem(db *sql.DB, item models.Item) error {
	_, err := db.Exec("INSERT INTO items(name) values($1)", item.Name)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItem delete item with id
func DeleteItem(db *sql.DB, id int) error {
	_, err := db.Exec("delete from items where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
