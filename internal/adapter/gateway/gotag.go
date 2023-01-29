package gateway

import (
	"context"
	"fmt"

	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/domain/model/gotag"
	"github.com/takokun778/gotagnews/internal/domain/repository"
	"github.com/takokun778/gotagnews/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ repository.Gotag = (*Gotag)(nil)

const (
	DatabaseName   = "gotagdatabase"
	CollectionName = "gotagcollection"
	UniqueKey      = "tag"
)

type Gotag struct {
	collection *mongo.Collection
}

type coll struct {
	ID  primitive.ObjectID `json:"id" bson:"_id"`
	Tag string             `json:"tag" bson:"tag"`
}

func NewGotag(
	client *DB,
) *Gotag {
	return &Gotag{
		collection: client.Database(DatabaseName).Collection(CollectionName),
	}
}

func (gt *Gotag) FindAll(ctx context.Context) (model.GotagList, error) {
	cur, err := gt.collection.Find(ctx, bson.M{})
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to find all gotag", log.ErrorField(err))

		return nil, fmt.Errorf("failed to find all gotag %w", err)
	}

	list := make(model.GotagList, 0)

	for cur.Next(ctx) {
		var col coll
		if err := cur.Decode(&col); err != nil {
			log.GetLogCtx(ctx).Warn("failed to decode gotag", log.ErrorField(err))

			return nil, fmt.Errorf("failed to decode gotag %w", err)
		}

		list = append(list, model.Gotag{
			ID: gotag.ID(col.Tag),
		})
	}

	return list, nil
}

func (gt *Gotag) SaveAll(ctx context.Context, list model.GotagList) error {
	if len(list) == 0 {
		log.GetLogCtx(ctx).Info("no gotag to save")

		return nil
	}

	docs := make([]interface{}, len(list))

	for i, item := range list {
		docs[i] = bson.D{{Key: UniqueKey, Value: item.ID.String()}}
	}

	if _, err := gt.collection.InsertMany(ctx, docs); err != nil {
		log.GetLogCtx(ctx).Warn("failed to save all gotag", log.ErrorField(err))

		return fmt.Errorf("failed to save all gotag %w", err)
	}

	return nil
}
