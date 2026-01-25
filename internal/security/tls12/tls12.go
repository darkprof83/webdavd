package tls12

import (
	"crypto/tls"
	"net/http"

	"github.com/darkprof83/webdavd/internal/config"
)

type TLS12 struct{}

func (s *TLS12) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
}

func (s *TLS12) Server(router http.Handler, cfg *config.Config) *http.Server {
	tlscfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
		TLSConfig:    tlscfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	return srv
}

func (s *TLS12) ListenAndServe(srv *http.Server, cfg *config.Config) error {
	return srv.ListenAndServeTLS(cfg.Cert, cfg.Key)
}
