package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

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
	CloseChan = make(chan os.Signal)
	signal.Notify(CloseChan, os.Interrupt)
	go listenForClose()
	FileName = fileName
	router := mux.NewRouter()
	router.HandleFunc(EndPoint, MainHandler)
	log.Println("starting server")
	err := http.ListenAndServe(":"+Port, router)
	if err != nil {
		log.Fatal("Major error")
	}
}

func StartServerWithConfig(configFileName string) {
	CloseChan = make(chan os.Signal)
	signal.Notify(CloseChan, os.Interrupt)
	go listenForClose()
	configFileContent, err := ReadFile(configFileName)
	if err != nil {
		log.Println("Error during read", err)
		return
	}
	var config interface{}
	err = json.Unmarshal(configFileContent, &config)
	if err != nil {
		log.Println("Error during unmarshal", err)
		return
	}
	ParsedConfig :=ReadConfig(config)
	router := mux.NewRouter()
	for _ , value := range ParsedConfig.Endpoints{
		router.HandleFunc(value.Url, GenerateHandler(value))
	}
	log.Println("server started..")
	err = http.ListenAndServe(":"+ParsedConfig.Global.Port , router)
	if err != nil{
		fmt.Println(" webserver error")
	}
}

func listenForClose() {
	for {
		select {
		case <-CloseChan:
			fmt.Println("Gracefully shutting down...")
			os.Exit(1)
		}
	}
}
