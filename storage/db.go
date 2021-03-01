package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
)

// InitDataBase postgres init
func InitDataBase() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("storage.postgres.username"),
		viper.GetString("storage.postgres.password"),
		viper.GetString("storage.postgres.host"),
		viper.GetString("storage.postgres.port"),
		viper.GetString("storage.postgres.dbname"),
	)

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
