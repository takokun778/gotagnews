package external

import (
	"context"

	"github.com/takokun778/gotagnews/internal/domain/model"
)

type Gotag interface {
	Notice(context.Context, model.GotagList) error
}
