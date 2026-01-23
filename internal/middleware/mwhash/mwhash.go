package mwhash

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/darkprof83/webdavd/internal/hash"
)

func New(cfgsalt string, cfguser string, cfghash string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok {
				basicAuthFailed(w, "Restricted")
				return
			}
			isValidUserName := user == cfguser
			// TODO: move spritf in hash
			tested := fmt.Sprintf("%x", hash.Get256(cfgsalt, pass))
			isValidPassHash := strings.Compare(cfghash, tested) == 0
			if !isValidUserName || !isValidPassHash {
				basicAuthFailed(w, "Need authorized!")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func basicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
