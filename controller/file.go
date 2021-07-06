package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/docker-golang-mysql/model"
	"github.com/gorilla/mux"
)

const urlImage = "http://localhost:8080/file/readfile/"

func Upload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("newImage")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	// tao ra 1 file trong
	tempFile, err := ioutil.TempFile("upload", "*.png")
	if err != nil {
		fmt.Println(err)
	}
	// doc file upload
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// viet vao file trong
	tempFile.Write(fileBytes)
	// lua url
	nameFile := tempFile.Name()
	nameImage := nameFile[7:]
	url := urlImage + nameImage
	fmt.Fprint(w, url)
	// fmt.Print(url)
}

//(`upload/upload-780556805.png`)

func ReadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	image := vars["image"]
	// fmt.Fprint(w, image)
	url := "upload/" + image
	data, err := ioutil.ReadFile(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(data))
}

func WriteCSV(w http.ResponseWriter, r *http.Request) {
	var data []model.Person
	data = model.GetPersonAPI()
	// var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}

	fmt.Print(data)

	file, err := os.Create("result.csv")
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data {
		data, _ := json.Marshal(value)
		valuedata := []string{string(data)}
		err := writer.Write(valuedata)
		if err != nil {
			fmt.Println(err)
		}
	}
}
