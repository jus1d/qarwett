package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	EnvLocal       = "local"
	EnvDevelopment = "dev"
	EnvProduction  = "prod"
)

type Config struct {
	Env       string    `yaml:"env" env-required:"true"`
	Telegram  Telegram  `yaml:"telegram"`
	ICalendar ICalendar `yaml:"icalendar"`
	Postgres  Postgres  `yaml:"postgres"`
}

type Telegram struct {
	Token string `yaml:"token"`
}

type ICalendar struct {
	Updater Updater `yaml:"updater"`
	Server  Server  `yaml:"server"`
}

type Updater struct {
	WeeksToTrack int `yaml:"weeks_to_track"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	ModeSSL  string `yaml:"sslmode"`
}

// MustLoad loads config to a new Config instance and return it.
func MustLoad() *Config {
	_ = godotenv.Load()

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatalf("missed CONFIG_PATH parameter")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at: %s", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &config
}
