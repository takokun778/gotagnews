package repository

import (
	"context"

	"github.com/takokun778/gotagnews/internal/domain/model"
)

type GitHub interface {
	FindAll(context.Context) (model.GotagList, error)
}
