package notifier

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Channel struct {
	Client *linebot.Client
}

type ChannelFactory interface {
	Of(string, string) (*Channel, error)
}
