package main

import (
	"context"
	"os"
	"time"

	"github.com/takokun778/gotagnews/internal/adapter/gateway"
	"github.com/takokun778/gotagnews/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout = 60 * time.Second
	value   = "migration"
)

func main() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	uri := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Log().Panic("failed to connect to MongoDB: ", log.ErrorField(err))
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, nil); err != nil {
		log.Log().Panic("failed to ping to MongoDB: ", log.ErrorField(err))
	}

	dbname := gateway.DatabaseName

	collectionname := gateway.CollectionName

	coll := client.Database(dbname).Collection(collectionname)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: gateway.UniqueKey, Value: -1}},
		Options: options.Index().SetUnique(true),
	}

	if _, err = coll.Indexes().CreateOne(context.TODO(), indexModel); err != nil {
		log.Log().Panic("failed to create index: ", log.ErrorField(err))
	}

	result, err := coll.InsertOne(ctx, bson.D{{Key: gateway.UniqueKey, Value: value}})
	if err != nil {
		log.Log().Panic("failed to insert collection: ", log.ErrorField(err))
	}

	log.Log().Sugar().Infof("inserted %+v collection", result)

	filter := bson.D{{Key: gateway.UniqueKey, Value: value}}

	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		log.Log().Panic("failed to delete collection: ", log.ErrorField(err))
	}
}
