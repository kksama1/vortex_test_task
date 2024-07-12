package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

func LoadConfig[Config DatabaseConfig]() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	var cfg Config
	if err = envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("cannot process env : %v", err)
	} else {
		log.Println("config initialized")
	}
	return &cfg, nil
}
