package mock

import (
	"context"
	"testing"

	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/domain/repository"
)

var _ repository.GitHub = (*GitHub)(nil)

type GitHub struct {
	T          *testing.T
	List       model.GotagList
	ErrFindAll error
}

func (gh *GitHub) FindAll(_ context.Context) (model.GotagList, error) {
	gh.T.Helper()

	return gh.List, gh.ErrFindAll
}
