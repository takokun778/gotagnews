package line

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/takokun778/gotagnews/internal/adapter/notifier"
	"github.com/takokun778/gotagnews/pkg/log"
)

var _ notifier.ChannelFactory = (*Client)(nil)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Of(secret, token string) (*notifier.Channel, error) {
	client, err := linebot.New(secret, token)
	if err != nil {
		log.Log().Warn("failed to create line client", log.ErrorField(err))

		return nil, fmt.Errorf("failed to create line client: %w", err)
	}

	if _, err := client.GetBotInfo().Do(); err != nil {
		log.Log().Warn("failed to get bot info", log.ErrorField(err))

		return nil, fmt.Errorf("failed to get bot info: %w", err)
	}

	return &notifier.Channel{
		Client: client,
	}, nil
}
