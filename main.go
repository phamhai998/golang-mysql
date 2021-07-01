package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/docker-golang-mysql/controller"
	"github.com/docker-golang-mysql/model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("web runing ....")
	model.ConnectDB()
	router := mux.NewRouter()

	router.HandleFunc("/person", controller.GetPerson).Methods("GET")
	router.HandleFunc("/person/insert", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/person/search", controller.SearchPerson).Methods("POST")
	router.HandleFunc("/person/update", controller.UpdatePerson).Methods("POST")
	router.HandleFunc("/person/delete", controller.DeletePerson).Methods("POST")
	router.HandleFunc("/upload", controller.Upload).Methods("POST")
	router.HandleFunc("/html", controller.Html).Methods("GET")
	router.HandleFunc("/readfile", controller.ReadFile).Methods("GET")
	router.HandleFunc("/writefile", controller.WriteCSV).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "multipart/form-data"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
