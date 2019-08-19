package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var CloseChan chan os.Signal = make(chan os.Signal)

func MainHandler(res http.ResponseWriter, req *http.Request) {
	thisFile, err := os.Open(FileName)
	if err != nil {
		log.Println("opening error", err)
	}
	data, err := ioutil.ReadAll(thisFile)
	if err != nil {
		log.Println("reading error", err)
	}
	res.Header().Set("content-type", "application/json")
	res.Write(data)
}

func StartServer(fileName string) {
	FileName = fileName
	router := mux.NewRouter()
	router.HandleFunc(EndPoint, MainHandler)
	log.Println("starting server")
	err := http.ListenAndServe(":"+Port, router)
	if err != nil {
		log.Fatal("Major error")
	}
}
