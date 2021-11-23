package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(k string) string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return os.Getenv(k)
}
