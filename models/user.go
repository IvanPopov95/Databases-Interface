package models

// User main struct
type User struct {
	ID   int    `bson:"id" db:"id"`
	Name string `bson:"name" db:"name"`
}
