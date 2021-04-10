package repository

import (
	"context"
	"projectttt/databases"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Collection() *mongo.Collection
	Create(ctx context.Context, item interface{}) error
	GetOneElem(ctx context.Context, elem interface{}, id interface{}, field string) error
	GetElems(ctx context.Context, elements interface{}) error
	DeleteElem(ctx context.Context, id interface{}) error
}

type repo struct {
	database   string
	collection string
}

func (r *repo) Collection() *mongo.Collection {
	return databases.GetClient().Database(r.database).Collection(r.collection)
}

// Create adding element to mongodb
func (r *repo) Create(ctx context.Context, item interface{}) error {
	_, err := r.Collection().InsertOne(ctx, item)
	return err
}

// GetOneElem get one element from mongodb
func (r *repo) GetOneElem(ctx context.Context, elem interface{}, id interface{}, field string) error {
	res := r.Collection().FindOne(ctx, bson.M{field: id})
	return res.Decode(elem)
}

// GetElems return elements from mongodb
func (r *repo) GetElems(ctx context.Context, elements interface{}) error {
	cursor, err := r.Collection().Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, elements)
}

// DeleteElem delete one element with id
func (r *repo) DeleteElem(ctx context.Context, id interface{}) error {
	_, err := r.Collection().DeleteOne(ctx, bson.M{"id": id})
	return err
}
