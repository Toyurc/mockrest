package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		// panic("file opening error")
		return nil, err
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

func GenerateHandler(e *EndpointObject) func(w http.ResponseWriter, r *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("content-type", "application/json")
		res.Write([]byte(e.Payload))
	}
}

func ReadConfig(configObject interface{}) *Config {
	PayloadData := configObject.(map[string]interface{})
	globals := PayloadData["global"].(map[string]interface{})
	// get port from global
	port := globals["port"].(string)
	var NewPayload string = ""
	endpoints := PayloadData["endpoints"].([]interface{})
	newEndPoints := []*EndpointObject{}
	for _, value := range endpoints {
		eachEndpoint := value.(map[string]interface{})
		eachUrl := eachEndpoint["url"].(string)
		if this, ok := eachEndpoint["payload"].([]interface{}); ok {
			data, err := json.Marshal(this)
			if err != nil {
				fmt.Println("error", err)
			}
			NewPayload = string(data)
		} else if this, ok := eachEndpoint["payload"].(map[string]interface{}); ok {
			data, err := json.Marshal(this)
			if err != nil {
				fmt.Println("error", err)
			}
			NewPayload = string(data)
		} else {
			panic("invalid payload")
		}
		newEndPoints = append(newEndPoints, &EndpointObject{
			Url:     eachUrl,
			Payload: NewPayload,
		})
	}

	return &Config{
		Global:    &Global{port},
		Endpoints: newEndPoints,
	}
}