package main

import (
	"fmt"
	"net/http"

	"github.com/docker-golang-mysql/controller"
	"github.com/docker-golang-mysql/model"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("web runing ....")
	model.ConnectDB()
	router := mux.NewRouter()

	router.HandleFunc("/person", controller.GetPerson).Methods("GET")
	router.HandleFunc("/person", controller.CreatePerson).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
