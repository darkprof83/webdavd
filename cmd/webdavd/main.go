package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/darkprof83/webdavd/internal/config"
	"github.com/darkprof83/webdavd/internal/logger"
	"github.com/darkprof83/webdavd/internal/middleware/mwhash"
	"github.com/darkprof83/webdavd/internal/middleware/mwlog"
	"github.com/darkprof83/webdavd/pkg/toattr"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/webdav"
)

const (
	configPath = "~/.config/webdavd/webdavd.yaml"
)

func init() {
	chi.RegisterMethod("PROPFIND")
	chi.RegisterMethod("PROPPATCH")
	chi.RegisterMethod("MKCOL")
	chi.RegisterMethod("COPY")
	chi.RegisterMethod("MOVE")
	chi.RegisterMethod("LOCK")
	chi.RegisterMethod("UNLOCK")
	chi.RegisterMethod("DELETE")
	chi.RegisterMethod("HEAD")
	chi.RegisterMethod("OPTIONS")
}

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

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mwlog.New(&log.Logger))
	router.Use(mwhash.New(cfg.Salt, cfg.Username, cfg.Passhash))
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Recoverer)

	fs := &webdav.Handler{
		Prefix:     cfg.Prefix,
		FileSystem: webdav.Dir(cfg.Dir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.ErrorContext(r.Context(), err.Error())
			} else {
				log.DebugContext(r.Context(), "webdav request")
			}
		},
	}
	router.Handle(cfg.Prefix, fs)

	log.Info("starting server",
		slog.String("address", cfg.Address),
		slog.String("prefix", cfg.Prefix),
		slog.String("dir", cfg.Dir),
	)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server", toattr.Err(err))
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HaltTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", toattr.Err(err))

		return
	}

	log.Info("server stopped")
}
