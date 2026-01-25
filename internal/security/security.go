package security

import (
	"net/http"

	"github.com/darkprof83/webdavd/internal/config"
	"github.com/darkprof83/webdavd/internal/security/none"
	"github.com/darkprof83/webdavd/internal/security/tls12"
)

type Securer interface {
	http.Handler
	Server(http.Handler, *config.Config) *http.Server
	ListenAndServe(*http.Server, *config.Config) error
}

var Security = map[string]Securer{
	"none":  &none.None{},
	"tls12": &tls12.TLS12{},
}
