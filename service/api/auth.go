package main

import (
	"net/http"
	"github.com/karlpokus/opentracing-lab/service/utils/logs"
)

func authenticate(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			logs.Stderr.Print("Auth header missing")
			http.Error(w, http.StatusText(401), 401)
			return
		}

		err := checkUser(authHeader)
		if err != nil {
			logs.Stderr.Print(err.Error())
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next.ServeHTTP(w, r)
	}
}
