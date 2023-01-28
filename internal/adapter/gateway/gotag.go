package gateway

import (
	"context"

	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/domain/model/gotag"
	"github.com/takokun778/gotagnews/internal/domain/repository"
	"github.com/takokun778/gotagnews/pkg/log"
)

var _ repository.Gotag = (*Gotag)(nil)

type Gotag struct{}

func NewGotag() *Gotag {
	return &Gotag{}
}

func (gt *Gotag) FindAll(ctx context.Context) (model.GotagList, error) {
	return model.GotagList{
		model.Gotag{ID: gotag.ID("a")},
		model.Gotag{ID: gotag.ID("b")},
	}, nil
}

func (gt *Gotag) SaveAll(ctx context.Context, list model.GotagList) error {
	logger := log.GetLogCtx(ctx)

	for _, item := range list {
		logger.Sugar().Infof("save: %s", item.ID)
	}

	return nil
}
