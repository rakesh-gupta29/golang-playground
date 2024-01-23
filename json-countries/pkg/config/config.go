package config

import (
  "log"

  "github.com/caarlos0/env"
  "github.com/joho/godotenv"
)
type AppConfig struct {
  Address string `env:"address" envDefault:"localhost:4000"`
  Mode string `env:"MODE"`
}

func LoadConfig()( *AppConfig, error) {
  err := godotenv.Load("config.env")

  if  err != nil {
    log.Println("Error Loading ENV Variables.", err)
  }
  cfg := AppConfig{}
  err = env.Parse(&cfg) 
  if err != nil {
    return nil, err
  }
  return &cfg, nil
}
