package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/docker-golang-mysql/model"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("newImage")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("upload", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
}

func ReadFile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("readfile.txt")
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
