package databases

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// InitMongoDataBase connect to mongo db
func InitMongoDataBase() error {
	connSettings := fmt.Sprintf("mongodb://%s:%s",
		viper.GetString("storage.mongo.host"),
		viper.GetString("storage.mongo.port"))

	opts := options.Client().ApplyURI(connSettings)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	client = c

	return client.Ping(ctx, readpref.Primary())

}

func GetClient() *mongo.Client {
	return client
}
