package requestLogger

import (
	"net/http"
	"log"
)

func Log(logger *log.Logger, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s %s", r.Method, r.URL)
		h.ServeHTTP(w, r)
	}
}
