package config

type Config struct{}

var config Config //nolint:gochecknoglobals

func Init() {
	config = Config{}
}

func Get() Config {
	return config
}
