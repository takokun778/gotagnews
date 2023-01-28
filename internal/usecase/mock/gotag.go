package mock

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/takokun778/gotagnews/internal/domain/external"
	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/domain/repository"
)

var (
	_ repository.Gotag = (*Gotag)(nil)
	_ external.Gotag   = (*Gotag)(nil)
)

type Gotag struct {
	T          *testing.T
	List       model.GotagList
	Want       model.GotagList
	ErrSaveAll error
	ErrFindAll error
	ErrNotice  error
}

func (g *Gotag) SaveAll(ctx context.Context, list model.GotagList) error {
	g.T.Helper()

	return g.ErrSaveAll
}

func (g *Gotag) FindAll(ctx context.Context) (model.GotagList, error) {
	g.T.Helper()

	return g.List, g.ErrFindAll
}

func (g *Gotag) Notice(ctx context.Context, list model.GotagList) error {
	g.T.Helper()

	if g.ErrNotice != nil {
		return g.ErrNotice
	}

	sort.SliceStable(list, func(i, j int) bool { return list[i].ID.String() < list[j].ID.String() })

	sort.SliceIsSorted(g.Want, func(i, j int) bool { return g.Want[i].ID.String() < g.Want[j].ID.String() })

	if !reflect.DeepEqual(list, g.Want) {
		g.T.Errorf("GotagNotice = %v, want %v", list, g.Want)

		return nil
	}

	return nil
}
