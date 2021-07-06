package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/docker-golang-mysql/model"
	"github.com/docker-golang-mysql/validate"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	ListPerson := model.GetPersonAPI()
	json.NewEncoder(w).Encode(ListPerson)
	// fmt.Fprint(w, ListPerson[0])
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		log.Fatalln(err)
	}
	if err := validate.ValidateInputPerson(&person); err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprint(w, "successfully")
		model.CreatePersonAPI(person)
	}
	// fmt.Print(person)
}

func SearchPerson(w http.ResponseWriter, r *http.Request) {
	var s model.S
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		log.Fatalln(err)
	}
	listPerson := model.SearchPersonAPI(s)
	json.NewEncoder(w).Encode(listPerson)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatalln(err)
	}
	model.UpdatePersonAPI(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	var id model.ID
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Print(id)
	model.DeletePersonAPI(id)
}

type NewsAggPage struct {
	Title string
	News  string
}

func Html(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "test", News: "hai"}
	t, _ := template.ParseFiles("test.html")
	t.Execute(w, p)
}
