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

// GetUsersList get all users from db
func GetUsersList(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var item models.User
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, item)
	}
	return users, nil
}

// GetUserWithID get one user with id
func GetUserWithID(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// AddUser adding one user to db
func AddUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users(name) values($1)", user.Name)
	return err
}

// DeleteUser delete user with id
func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
