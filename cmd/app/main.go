package main

import (
	"context"

	"github.com/takokun778/gotagnews/internal/adapter/controller"
	"github.com/takokun778/gotagnews/internal/adapter/gateway"
	"github.com/takokun778/gotagnews/internal/adapter/notifier"
	"github.com/takokun778/gotagnews/internal/usecase/interactor"
	"github.com/takokun778/gotagnews/pkg/log"
)

func main() {
	gotagRepository := gateway.NewGotag()

	githubRepository := gateway.NewGitHub()

	gotagExternal := notifier.NewGotag()

	usecase := interactor.NewGotagNotice(gotagRepository, githubRepository, gotagExternal)

	cmd := controller.NewGotag(usecase)

	if err := cmd.Cmd(context.Background()); err != nil {
		log.Log().Panic("failed to run command", log.ErrorField(err))
	}
}
