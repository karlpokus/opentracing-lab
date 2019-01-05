package main

import (
	"net/http"
	"log"
	"context"
	"time"
	"os"
	router "github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/karlpokus/opentracing-lab/service/utils/requestLogger"
)

type server struct {
	router *router.Router
	stdout *log.Logger
	stderr *log.Logger
	pets *mongo.Collection
}

func (s *server) applyRoutes() {
	s.router.HandlerFunc("GET", "/pet/:name", s.findOnePetHandler)
	s.router.HandlerFunc("GET", "/pets", s.findAllPetsHandler)
	s.router.HandlerFunc("POST", "/pets/add", s.addOnePetHandler)
}

func (s *server) findOnePetHandler(w http.ResponseWriter, r *http.Request) {
	petName := router.ParamsFromContext(r.Context())[0].Value
	pet, err := findOnePet(s.pets, petName)
	if err != nil {
		s.stderr.Print(err.Error())
		http.Error(w, http.StatusText(404), 404)
		return
	}

	s.stdout.Println("ok")
	w.Header().Set("Content-Type", "application/json")
	w.Write(pet)
}

func (s *server) findAllPetsHandler(w http.ResponseWriter, r *http.Request) {
	pets, err := findAllPets(s.pets)
	if err != nil {
		s.stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}

	s.stdout.Println("ok")
	w.Header().Set("Content-Type", "application/json")
	w.Write(pets)
}

func (s *server) addOnePetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if err := addOnePet(s.pets, r.Body); err != nil {
		s.stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}

	s.stdout.Println("ok")
	w.Write([]byte("ok"))
}

func main() {
	s := &server{
		router: router.New(),
		stdout: log.New(os.Stdout, "", 0),
		stderr: log.New(os.Stderr, "", 0),
	}
	s.applyRoutes()

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	mongoConnString := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, mongoConnString)
	if err != nil {
		s.stderr.Fatal(err.Error())
	}

	s.pets = client.Database("pets").Collection("pets")
	s.stdout.Printf("connected to %s", mongoConnString)
	s.stdout.Print("listening on port 9113")
	s.stderr.Fatal(http.ListenAndServe(":9113", requestLogger.Log(s.stdout, s.router)))
}
