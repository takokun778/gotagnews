package notifier

import (
	"context"

	"github.com/takokun778/gotagnews/internal/domain/external"
	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/pkg/log"
)

var _ external.Gotag = (*Gotag)(nil)

type Gotag struct{}

func NewGotag() *Gotag {
	return &Gotag{}
}

func (gt *Gotag) Notice(ctx context.Context, list model.GotagList) error {
	logger := log.GetLogCtx(ctx)

	for _, item := range list {
		logger.Sugar().Infof("notice: %s", item.ID)
	}

	return nil
}
