// Package config 配置包
package config

import (
	"github.com/joho/godotenv"

	"fmt"
	"os"
)

func init() {
	err := godotenv.Load(".upcloud.env")
	if err != nil {
		fmt.Println(".upcloud.env the configuration file does not exist")
		os.Exit(0)
	}
}
