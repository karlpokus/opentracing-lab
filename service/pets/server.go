package main

import (
	"net/http"
	"log"
	"context"
	"time"
	"os"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/karlpokus/opentracing-lab/service/utils/requestLogger"
)

var (
	stdout *log.Logger = log.New(os.Stdout, "", 0)
	stderr *log.Logger = log.New(os.Stderr, "", 0)
)

type Env struct {
	pets *mongo.Collection
}

func (env *Env) findPet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
  	http.Error(w, http.StatusText(405), 405)
    return
  }

	petName := r.URL.Query().Get("name")
	pet, err := findPet(env.pets, petName)
	if err != nil {
		stderr.Print(err.Error())
		http.Error(w, http.StatusText(404), 404)
		return
	}

	stdout.Print("success")
	// w.Header().Set("Content-Type", "application/json")
	w.Write(pet)
}

func (env *Env) findAllPets(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
  	http.Error(w, http.StatusText(405), 405)
    return
  }

	pets, err := findAllPets(env.pets)
	if err != nil {
		stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}

	stdout.Print("success")
	w.Write(pets)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	mongoConnString := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, mongoConnString)
	if err != nil {
		stderr.Fatal(err.Error())
	}
	stdout.Printf("connected to %s", mongoConnString)

	env := &Env{pets: client.Database("pets").Collection("pets")}

	http.HandleFunc("/api/pets", env.findAllPets);
	http.HandleFunc("/api/pet", env.findPet);
	stdout.Print("listening on port 9113")
	stderr.Fatal(http.ListenAndServe(":9113", requestLogger.Log(stdout, http.DefaultServeMux)))
}
