package models

// Item main struct
type Item struct {
	ID   int    `bson:"id" db:"id"`
	Name string `bson:"name" db:"name"`
}
