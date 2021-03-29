package dbmongo

import (
	"context"
	"fmt"
	"projectttt/models"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// InitDataBase connect to mongo db
func InitDataBase() (*mongo.Collection, error) {
	connSettings := fmt.Sprintf("mongodb://%s:%s",
		viper.GetString("storage.mongo.host"),
		viper.GetString("storage.mongo.port"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connSettings))
	if err != nil {
		return nil, err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	collection := client.Database("testing").Collection("numbers")
	return collection, nil
}

// AddToDb adding element to mongodb
func AddToDb(collection *mongo.Collection, doc interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func GetItemWithID get one item from mongo db
func GetItemWithID(collection *mongo.Collection, id int) (*models.Item, error) {
	var item models.Item
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func DeleteItem(collection *mongo.Collection, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.D{{"id", id}})
	if err != nil {
		return err
	}
	return nil
}
