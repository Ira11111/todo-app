package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env     string     `yaml:"env"`
	Storage string     `yaml:"storage_path"`
	Server  HTTPConfig `yaml:"http_server"`
}

type HTTPConfig struct {
	Addr        string `yaml:"address"`
	Timeout     int    `yaml:"timeout"`
	IdleTimeout int    `yaml:"idle_timeout"`
}

// MustLoad не возвращает ошибку а паникует
func MustLoad() *Config {
	path := fetchConfig()
	if path == "" {
		panic("config file path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}

func fetchConfig() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	res := os.Getenv("CONFIG_FILE")
	return res
}
