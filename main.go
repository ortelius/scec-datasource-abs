package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ortelius/scec-datasource-abs/handler"
)

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/upload", handler.Create).Methods("POST")
	myRouter.HandleFunc("/status", handler.GetStatus).Methods("GET")

	handler.SetSwagger(myRouter)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
