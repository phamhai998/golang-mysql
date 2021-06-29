package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/docker-golang-mysql/model"
)

type Preson struct {
	ID    int
	Name  string
	Email string
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	ListPerson := model.GetPersonAPI()
	json.NewEncoder(w).Encode(ListPerson)
	// fmt.Fprint(w, ListPerson[0])
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatalln(err)
	}
	model.CreatePersonAPI(person)
}
