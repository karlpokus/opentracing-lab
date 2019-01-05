package main

import (
	"net/http"
	"log"
	"os"
	"github.com/karlpokus/opentracing-lab/service/utils/requestLogger"
	router "github.com/julienschmidt/httprouter"
)

type server struct {
	router *router.Router
	stdout *log.Logger
	stderr *log.Logger
}

func (s *server) applyRoutes() {
	s.router.Handler("GET", "/api/pet/:name", s.authenticate(s.findOnePet()))
	s.router.Handler("GET", "/api/pets", s.authenticate(s.findAllPets()))
	s.router.Handler("POST", "/api/pets/add", s.authenticate(s.addOnePet()))
}

func (s *server) authenticate(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			s.stderr.Print("Auth header missing")
			http.Error(w, http.StatusText(401), 401)
			return
		}

		err := checkUser(authHeader)
		if err != nil {
			s.stderr.Print(err.Error())
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (s *server) findOnePet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := router.ParamsFromContext(r.Context())[0].Value
		path := "/pet/" + name
		pet, err := httpReq("GET", path, nil)
		if err != nil {
			s.stderr.Print(err.Error())
			http.Error(w, http.StatusText(404), 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(pet)
	}
}

func (s *server) findAllPets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pets, err := httpReq("GET", "/pets", nil)
		if err != nil {
			s.stderr.Print(err.Error())
			http.Error(w, http.StatusText(400), 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(pets)
	}
}

func (s *server) addOnePet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			errMsg := "POST body missing"
			s.stderr.Print(errMsg)
			http.Error(w, errMsg, 400)
			return
		}

		res, err := httpReq("POST", "/pets/add", r.Body)
		if err != nil {
			s.stderr.Print(err.Error())
			http.Error(w, http.StatusText(400), 400)
			return
		}

		w.Write(res)
	}
}

func main() {
	s := &server{
		router: router.New(),
		stdout: log.New(os.Stdout, "", 0),
		stderr: log.New(os.Stderr, "", 0),
	}
	s.applyRoutes()

	s.stdout.Print("listening on port 9111")
	s.stderr.Fatal(http.ListenAndServe(":9111", requestLogger.Log(s.stdout, s.router)))
}
