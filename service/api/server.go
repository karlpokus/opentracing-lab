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
			errMsg := "Queryparam name missing"
			logs.Stderr.Print(errMsg)
			http.Error(w, errMsg, 400)
			return
		}

		pet, err := findOnePet(petName)
		if err != nil {
			logs.Stderr.Print(err.Error())
			http.Error(w, http.StatusText(404), 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(pet)
})

var findAllPetsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	pets, err := findAllPets()
	if err != nil {
		logs.Stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(pets)
})

var getPetTypesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	/*types, err := getPetTypes()
	if err != nil {
		logs.Stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(types)*/
	w.Write([]byte("hello from getPetTypesHandler"))
})

var addOnePetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		errMsg := "POST body missing"
		logs.Stderr.Print(errMsg)
		http.Error(w, errMsg, 400)
		return
	}

	res, err := addOnePet(r.Body)
	if err != nil {
		logs.Stderr.Print(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Write(res)
})

// init(env *Env) is a http.Handler
// it recieves an env and creates a ctx
// the rest of the middleware are custom handler i.e ctx, env, w, r

func main() {
	router := http.NewServeMux()
	router.Handle("/api/pet", methodAllowed.GET(logs.Stderr, authenticate(findOnePetHandler)))
	router.Handle("/api/pets", methodAllowed.GET(logs.Stderr, authenticate(findAllPetsHandler)))
	router.Handle("/api/pets/types", methodAllowed.GET(logs.Stderr, authenticate(getPetTypesHandler)))
	router.Handle("/api/pet/add", methodAllowed.POST(logs.Stderr, authenticate(addOnePetHandler)))

	logs.Stdout.Print("listening on port 9111")
	logs.Stderr.Fatal(http.ListenAndServe(":9111", requestLogger.Log(logs.Stdout, router)))
}
