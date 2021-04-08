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
	collection := client.Database("maindb").Collection("users")
	return collection, nil
}

// AddUser adding element to mongodb
func AddUser(collection *mongo.Collection, item models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserWithID get one user from mongodb
func GetUserWithID(collection *mongo.Collection, id int) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser delete user with id
func DeleteUser(collection *mongo.Collection, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.D{{"id", id}})
	return err
}

// GetUsersList return all users from database
func GetItemsList(collection *mongo.Collection) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var users []models.User
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
