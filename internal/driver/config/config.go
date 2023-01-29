package config

import "os"

type Config struct {
	MongoDBURI string
}

var config Config //nolint:gochecknoglobals

func Init() {
	config = Config{
		MongoDBURI: os.Getenv("MONGODB_URI"),
	}
}

func Get() Config {
	return config
}
