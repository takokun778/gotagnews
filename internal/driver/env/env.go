package env

import (
	"os"

	"github.com/takokun778/gotagnews/pkg/log"
)

type Env string

const (
	prod  Env = "prod"
	prev  Env = "prev"
	dev   Env = "dev"
	local Env = "local"
	empty Env = ""
)

var env Env //nolint:gochecknoglobals

func Init() {
	e := Env(os.Getenv("ENV"))

	log.Log().Sugar().Infof("environment: %s", e)

	env = e
}

func Get() Env {
	return env
}

func (e Env) String() string {
	return string(e)
}

func (e Env) IsProd() bool {
	return e == prod
}
