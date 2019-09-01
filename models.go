package main

type Global struct {
	Port string `json:"port"`
}

type EndpointObject struct {
	Url     string `json:"url"`
	Payload string `json:"payload"`
}

type Endpoints []*EndpointObject

type Config struct {
	Global *Global `json:"global"`
	Endpoints Endpoints `json:"endpoints"`
}
