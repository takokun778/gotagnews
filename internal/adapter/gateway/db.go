package gateway

import "go.mongodb.org/mongo-driver/mongo"

type DB struct {
	*mongo.Client
}

type DBFactory interface {
	Of(uri string) (*DB, error)
}
