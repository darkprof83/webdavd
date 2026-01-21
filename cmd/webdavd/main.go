package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/darkprof83/webdavd/internal/config"
	"github.com/darkprof83/webdavd/internal/logger"
)

const (
	configPath = "~/.config/webdavd/webdavd.yaml"
)

func main() {
	log := logger.New()
	log.Info("starting webdavd")

	var path string
	flag.StringVar(&path, "config", configPath, "path to the configuration file")
	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			log.Error(fmt.Sprintf("%v", r), slog.String("path", path))
			os.Exit(1)
		}
	}()

	cfg := config.New()
	cfg.MustLoad(path)
	if err := log.Setup(cfg.Env); err != nil {
		log.Error(err.Error(), slog.String("env", cfg.Env))
	}
	log.Info("the profile is used for logging", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
}
