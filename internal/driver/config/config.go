package config

import "os"

type Config struct {
	MongoDBURI string
	LINESecret string
	LINEToken  string
}

var config Config //nolint:gochecknoglobals

func Init() {
	config = Config{
		MongoDBURI: os.Getenv("MONGODB_URI"),
		LINESecret: os.Getenv("LINE_SECRET"),
		LINEToken:  os.Getenv("LINE_TOKEN"),
	}
}

func Get() Config {
	return config
}
