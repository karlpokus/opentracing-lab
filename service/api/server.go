package main

import (
	"net/http"
	"github.com/karlpokus/opentracing-lab/service/utils/requestLogger"
	"github.com/karlpokus/opentracing-lab/service/utils/methodAllowed"
	"github.com/karlpokus/opentracing-lab/service/utils/logs"
)

var findOnePetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		petName := r.FormValue("name")
		if len(petName) == 0 {
			logs.Stderr.Print("Queryparam name missing")
			http.Error(w, "Bad request. Queryparam name missing", 400)
			return
		}

		pet, err := findOnePet(petName)
		if err != nil {
			logs.Stderr.Print(err.Error())
			http.Error(w, http.StatusText(404), 404)
			return
		}
		w.Write(pet)
})

var findAllPetsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	pets, err := findAllPets()
	if err != nil {
		logs.Stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}
	w.Write(pets)
})

var getPetTypesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	types, err := getPetTypes()
	if err != nil {
		logs.Stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}
	w.Write(types)
})

func main() {
	router := http.NewServeMux()
	router.Handle("/api/pet", methodAllowed.GET(logs.Stderr, authenticate(findOnePetHandler)))
	router.Handle("/api/pets", methodAllowed.GET(logs.Stderr, authenticate(findAllPetsHandler)))
	router.Handle("/api/pets/types", methodAllowed.GET(logs.Stderr, authenticate(getPetTypesHandler)))

	logs.Stdout.Print("listening on port 9111")
	logs.Stderr.Fatal(http.ListenAndServe(":9111", requestLogger.Log(logs.Stdout, router)))
}
