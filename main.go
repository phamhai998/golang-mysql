package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/docker-golang-mysql/model"
	"github.com/docker-golang-mysql/routers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("web runing ....")
	model.ConnectDB()
	router := mux.NewRouter()

	routers.PersonRouters(router)
	routers.FileRouters(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "multipart/form-data"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
