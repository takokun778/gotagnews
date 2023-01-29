package mongo

import (
	"context"
	"fmt"

	"github.com/takokun778/gotagnews/internal/adapter/gateway"
	"github.com/takokun778/gotagnews/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ gateway.DBFactory = (*Client)(nil)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Of(uri string) (*gateway.DB, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Log().Warn("failed to connect to mongo", log.ErrorField(err))

		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Log().Warn("failed to ping mongo", log.ErrorField(err))

		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	log.Log().Info("success connected to mongo")

	return &gateway.DB{
		Client: client,
	}, nil
}
