package config

import (
	"go-blog/exception"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type ConfigImpl struct {
}

func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &ConfigImpl{}
}

func NewConfig() Config {
	return &ConfigImpl{}
}
