package config

import "os"

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (c *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func ProvideConfig() Config {
	return &configImpl{}
}
