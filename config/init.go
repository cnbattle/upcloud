package config

import (
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".upcloud.env")
	if err != nil {
		panic(err)
	}
}
