package models

// Item main struct
type Item struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
}
