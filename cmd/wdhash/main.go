package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/darkprof83/webdavd/internal/config"
	"github.com/darkprof83/webdavd/internal/hash"
	"github.com/darkprof83/webdavd/internal/logger"
	"github.com/darkprof83/webdavd/pkg/toattr"
)

const (
	configPath = "~/.config/webdavd/webdavd.yaml"
)

func main() {
	log := logger.New()

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a password: ")
	if input, err := reader.ReadString('\n'); err != nil {
		log.Error("error reading input", toattr.Err(err))
	} else {
		pass := strings.TrimSpace(input)
		fmt.Printf("%x\n", hash.Get256(cfg.Salt, pass))
		// log.Debug("computed", slog.String("pass", pass))
	}
}
