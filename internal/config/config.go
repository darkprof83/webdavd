package config

import (
	"os"

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
	Address  string `yaml:"address" env-required:"true"`
	Prefix   string `yaml:"prefix" env-default:"/GTD/"`
	Username string `yaml:"username" env-required:"true"`
	Salt     string `yaml:"salt" env-default:"test salt"`
	Passhash string `yaml:"passhash" env-required:"true"`
	Dir      string `yaml:"dir" env-required:"true"`
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
