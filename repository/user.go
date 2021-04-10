package repository

// User main struct
type User struct {
	ID   int    `bson:"id" db:"id"`
	Name string `bson:"name" db:"name"`
}

type userRepo struct {
	repo
}

func NewUserRepo() Repository {
	r := &userRepo{repo{database: "users", collection: "users"}}
	return Repository(r)
}
