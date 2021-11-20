package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var c *viper.Viper

func Init(env string) {
	c = viper.New()
	c.SetConfigFile("yaml")
	c.SetConfigName(env)
	path, err := filepath.Abs("config/environments")
	if err != nil {
		panic(err)
	}
	c.AddConfigPath(path)
	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Config() *viper.Viper {
	return c
}

func ResetPath() {
	path, err := filepath.Abs("../../")
	if err != nil {
		panic(err)
	}
	os.Chdir(path)
}
