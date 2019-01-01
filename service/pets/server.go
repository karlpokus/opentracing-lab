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

func (env *Env) findOnePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
  	http.Error(w, http.StatusText(405), 405)
    return
  }

	petName := r.URL.Query().Get("name")
	pet, err := findOnePet(env.pets, petName)
	if err != nil {
		stderr.Print(err.Error())
		http.Error(w, http.StatusText(404), 404)
		return
	}

	stdout.Println("ok")
	w.Header().Set("Content-Type", "application/json")
	w.Write(pet)
}

func (env *Env) findAllPetsHandler(w http.ResponseWriter, r *http.Request) {
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

	stdout.Println("ok")
	w.Header().Set("Content-Type", "application/json")
	w.Write(pets)
}

func (env *Env) addOnePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
  	http.Error(w, http.StatusText(405), 405)
    return
  }
	defer r.Body.Close()

	if err := addOnePet(env.pets, r.Body); err != nil {
		stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}

	stdout.Println("ok")
	w.Write([]byte("ok"))
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

	http.HandleFunc("/api/pet", env.findOnePetHandler);
	http.HandleFunc("/api/pets", env.findAllPetsHandler);
	http.HandleFunc("/api/pet/add", env.addOnePetHandler)

	stdout.Print("listening on port 9113")
	stderr.Fatal(http.ListenAndServe(":9113", requestLogger.Log(stdout, http.DefaultServeMux)))
}
