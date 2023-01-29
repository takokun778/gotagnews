package main

import (
	"context"

	"github.com/google/go-github/v50/github"
	"github.com/takokun778/gotagnews/internal/adapter/controller"
	"github.com/takokun778/gotagnews/internal/adapter/gateway"
	"github.com/takokun778/gotagnews/internal/adapter/notifier"
	"github.com/takokun778/gotagnews/internal/driver/config"
	"github.com/takokun778/gotagnews/internal/driver/mongo"
	"github.com/takokun778/gotagnews/internal/usecase/interactor"
	"github.com/takokun778/gotagnews/pkg/log"
)

func main() {
	ctx := log.SetLogCtx(context.Background())

	config.Init()

	db, err := mongo.NewClient().Of(config.Get().MongoDBURI)
	if err != nil {
		log.GetLogCtx(ctx).Panic("failed to connect to mongo", log.ErrorField(err))
	}

	defer func() {
		if err := db.Disconnect(ctx); err != nil {
			log.GetLogCtx(ctx).Error("failed to disconnect from mongo db", log.ErrorField(err))
		}
	}()

	gotagRepository := gateway.NewGotag(db)

	githubRepository := gateway.NewGitHub(github.NewClient(nil))

	gotagExternal := notifier.NewGotag()

	usecase := interactor.NewGotagNotice(gotagRepository, githubRepository, gotagExternal)

	cmd := controller.NewGotag(usecase)

	if err := cmd.Cmd(ctx); err != nil {
		log.GetLogCtx(ctx).Panic("failed to run command", log.ErrorField(err))
	}
}
