package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env      string         `yaml:"env"`
	Server   HTTPConfig     `yaml:"http_server"`
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

type HTTPConfig struct {
	Port        string `yaml:"port"`
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
