package notifier

import (
	"context"
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/takokun778/gotagnews/internal/domain/external"
	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/pkg/log"
)

var _ external.Gotag = (*Gotag)(nil)

type Gotag struct {
	channel *Channel
}

func NewGotag(
	channel *Channel,
) *Gotag {
	return &Gotag{
		channel: channel,
	}
}

func (gt *Gotag) Notice(ctx context.Context, list model.GotagList) error {
	logger := log.GetLogCtx(ctx)

	for _, item := range list {
		logger.Sugar().Infof("notice: %s", item.ID)

		content := linebot.NewTextMessage(fmt.Sprintf(model.Message, item.ID))

		if _, err := gt.channel.Client.BroadcastMessage(content).Do(); err != nil {
			logger.Sugar().Warn("failed to notice %w tag", item.ID, log.ErrorField(err))

			return fmt.Errorf("failed to notice %w", err)
		}
	}

	return nil
}
