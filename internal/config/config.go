package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Log    `yaml:"log"`
	Server `yaml:"server"`
}

type Log struct {
	Env string `yaml:"env" env-default:"local"`
}

type Server struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	HaltTimeout time.Duration `yaml:"halt_timeout" env-default:"10s"`
	Prefix      string        `yaml:"prefix" env-default:"/"`
	Username    string        `yaml:"username" env-required:"true"`
	Salt        string        `yaml:"salt" env-default:"test salt"`
	Passhash    string        `yaml:"passhash" env-required:"true"`
	Dir         string        `yaml:"dir" env-required:"true"`
}

func New() *Config {
	return &Config{}
}

func (cfg *Config) MustLoad(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic("failed to read the configuration file")
	}
}
