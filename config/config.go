package config

import "github.com/spf13/viper"

var c *viper.Viper

func Init(path, env string) {
	c = viper.New()
	c.SetConfigFile("yaml")
	c.SetConfigName(env)
	c.AddConfigPath(path)
	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetConfig() *viper.Viper {
	return c
}
