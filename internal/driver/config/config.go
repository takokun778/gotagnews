package config

import "os"

type Config struct {
	MongoDBURI        string
	LINEChannelSecret string
	LINEChannelToken  string
}

var config Config //nolint:gochecknoglobals

func Init() {
	config = Config{
		MongoDBURI:        os.Getenv("MONGODB_URI"),
		LINEChannelSecret: os.Getenv("LINE_CHANNEL_SECRET"),
		LINEChannelToken:  os.Getenv("LINE_CHANNEL_TOKEN"),
	}
}

func Get() Config {
	return config
}
