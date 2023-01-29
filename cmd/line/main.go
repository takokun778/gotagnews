package main

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/driver/config"
	"github.com/takokun778/gotagnews/internal/driver/line"
	"github.com/takokun778/gotagnews/pkg/log"
)

func main() {
	config.Init()

	channel, err := line.NewClient().Of(config.Get().LINEChannelSecret, config.Get().LINEChannelToken)
	if err != nil {
		log.Log().Panic("failed to connect to line", log.ErrorField(err))
	}

	log.Log().Info("connected to line")

	res, err := channel.Client.GetBotInfo().Do()
	if err != nil {
		log.Log().Panic("failed to get bot info", log.ErrorField(err))
	}

	log.Log().Sugar().Infof("bot info: %+v", res)

	content := linebot.NewTextMessage(fmt.Sprintf(model.Message, "go1.19.5"))

	if _, err := channel.Client.BroadcastMessage(content).Do(); err != nil {
		log.Log().Panic("failed to send message", log.ErrorField(err))
	}
}
