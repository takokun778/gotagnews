package gateway

import (
	"context"
	"fmt"

	"github.com/google/go-github/v50/github"
	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/domain/model/gotag"
	"github.com/takokun778/gotagnews/internal/domain/repository"
	"github.com/takokun778/gotagnews/pkg/log"
)

var _ repository.GitHub = (*GitHub)(nil)

const (
	perPage = 100
	page    = 1
)

type GitHub struct {
	*github.Client
}

func NewGitHub(
	client *github.Client,
) *GitHub {
	return &GitHub{
		Client: client,
	}
}

func (gh *GitHub) FindAll(ctx context.Context) (model.GotagList, error) {
	var res []model.Gotag

	opts := &github.ListOptions{
		PerPage: perPage,
		Page:    page,
	}

	for {
		tags, _, err := gh.Client.Repositories.ListTags(ctx, model.Owner, model.Repository, opts)
		if err != nil {
			log.GetLogCtx(ctx).Warn("failed to list tags", log.ErrorField(err))

			return nil, fmt.Errorf("failed to list tags: %w", err)
		}

		for _, tag := range tags {
			id, err := gotag.NewID(tag.GetName())
			if err != nil {
				log.GetLogCtx(ctx).Warn("failed new id", log.ErrorField(err))

				return nil, fmt.Errorf("failed new id: %w", err)
			}

			res = append(res, model.Gotag{
				ID: id,
			})
		}

		if len(tags) < perPage {
			break
		}

		opts.Page++
	}

	return res, nil
}
