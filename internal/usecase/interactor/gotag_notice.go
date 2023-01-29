package interactor

import (
	"context"
	"fmt"

	"github.com/takokun778/gotagnews/internal/domain/external"
	"github.com/takokun778/gotagnews/internal/domain/repository"
	"github.com/takokun778/gotagnews/internal/usecase/port"
	"github.com/takokun778/gotagnews/pkg/log"
)

var _ port.GotagNoticeUsecase = (*GotagNotice)(nil)

type GotagNotice struct {
	gotagRepository  repository.Gotag
	githubRepository repository.GitHub
	external         external.Gotag
}

func NewGotagNotice(
	gotagRepository repository.Gotag,
	githubRepository repository.GitHub,
	external external.Gotag,
) *GotagNotice {
	return &GotagNotice{
		gotagRepository:  gotagRepository,
		githubRepository: githubRepository,
		external:         external,
	}
}

func (gni *GotagNotice) Execute(
	ctx context.Context,
	input port.GotagNoticeInput,
) (port.GotagNoticeOutput, error) {
	src, err := gni.gotagRepository.FindAll(ctx)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to find all gotag from gotag repository", log.ErrorField(err))

		return port.GotagNoticeOutput{}, fmt.Errorf("failed to find all gotag from gotag repository: %w", err)
	}

	dst, err := gni.githubRepository.FindAll(ctx)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to find all gotag from github repository", log.ErrorField(err))

		return port.GotagNoticeOutput{}, fmt.Errorf("failed to find all gotag from github repository: %w", err)
	}

	res := dst.Take(src)

	if len(res) == 0 {
		log.GetLogCtx(ctx).Info("no new gotag found")

		return port.GotagNoticeOutput{}, nil
	}

	if err := gni.gotagRepository.SaveAll(ctx, res); err != nil {
		log.GetLogCtx(ctx).Warn("failed to save all gotag from gotag repository", log.ErrorField(err))

		return port.GotagNoticeOutput{}, fmt.Errorf("failed to save all gotag from gotag repository: %w", err)
	}

	if err := gni.external.Notice(ctx, res); err != nil {
		log.GetLogCtx(ctx).Warn("failed to notice gotag", log.ErrorField(err))

		return port.GotagNoticeOutput{}, fmt.Errorf("failed to notice gotag: %w", err)
	}

	return port.GotagNoticeOutput{}, nil
}
