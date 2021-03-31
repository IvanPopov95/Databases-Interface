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

// AddItem adding element to mongodb
func AddItem(collection *mongo.Collection, item models.Item) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetItemWithID get one item from mongo db
func GetItemWithID(collection *mongo.Collection, id int) (*models.Item, error) {
	var item models.Item
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// DeleteItem delete item with id
func DeleteItem(collection *mongo.Collection, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.D{{"id", id}})
	if err != nil {
		return err
	}
	return nil
}

// GetItemsList return all items from database
func GetItemsList(collection *mongo.Collection) ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var items []models.Item
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var item models.Item
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
