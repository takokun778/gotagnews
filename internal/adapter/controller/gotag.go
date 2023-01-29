package controller

import (
	"context"
	"fmt"

	"github.com/takokun778/gotagnews/internal/usecase/port"
	"github.com/takokun778/gotagnews/pkg/log"
)

type Gotag struct {
	usecase port.GotagNoticeUsecase
}

func NewGotag(
	usecase port.GotagNoticeUsecase,
) *Gotag {
	return &Gotag{
		usecase: usecase,
	}
}

func (gt *Gotag) Cmd(ctx context.Context) error {
	if _, err := gt.usecase.Execute(ctx, port.GotagNoticeInput{}); err != nil {
		log.GetLogCtx(ctx).Warn("failed to execute gotag notice", log.ErrorField(err))

		return fmt.Errorf("failed to execute gotag notice: %w", err)
	}

	return nil
}
