package methodAllowed

import (
	"net/http"
	"log"
)

func GET(logger *log.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			logger.Print(http.StatusText(405))
	  	http.Error(w, http.StatusText(405), 405)
	    return
	  }
		h.ServeHTTP(w, r)
	})
}

func POST(logger *log.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			logger.Print(http.StatusText(405))
	  	http.Error(w, http.StatusText(405), 405)
	    return
	  }
		h.ServeHTTP(w, r)
	})
}

func Allow(logger *log.Logger, m string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != m {
			logger.Print(http.StatusText(405))
	  	http.Error(w, http.StatusText(405), 405)
	    return
	  }
		h.ServeHTTP(w, r)
	})
}
