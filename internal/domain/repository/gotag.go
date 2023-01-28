package repository

import (
	"context"

	"github.com/takokun778/gotagnews/internal/domain/model"
)

type Gotag interface {
	SaveAll(context.Context, model.GotagList) error
	FindAll(context.Context) (model.GotagList, error)
}
