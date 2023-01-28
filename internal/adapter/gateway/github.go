package gateway

import (
	"context"

	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/domain/model/gotag"
	"github.com/takokun778/gotagnews/internal/domain/repository"
)

var _ repository.GitHub = (*GitHub)(nil)

type GitHub struct{}

func NewGitHub() *GitHub {
	return &GitHub{}
}

func (gh *GitHub) FindAll(ctx context.Context) (model.GotagList, error) {
	return model.GotagList{
		model.Gotag{ID: gotag.ID("a")},
		model.Gotag{ID: gotag.ID("b")},
		model.Gotag{ID: gotag.ID("c")},
	}, nil
}
