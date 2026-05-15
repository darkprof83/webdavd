package none

import (
	"net/http"

	"github.com/darkprof83/webdavd/internal/config"
)

type None struct{}

func (s *None) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func (s *None) Server(router http.Handler, cfg *config.Config) *http.Server {
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return srv
}

func (s *None) ListenAndServe(srv *http.Server, cfg *config.Config) error {
	return http.ListenAndServe(cfg.Address, srv.Handler)
}
