package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Workers struct {
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/info", info).Methods("GET")
	router.HandleFunc("/wokers", getWorkers).Methods("GET")
	router.HandleFunc("/woker", registerWorker).Methods("POST")
	router.HandleFunc("/woker", upateWorker).Methods("PUT")

	err := http.ListenAndServe(":8089", router)
	if err != nil {
		log.Fatal("Start server err: ", err)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	// Redirect reqeust to workers
}

func getWorkers(w http.ResponseWriter, r *http.Request) {
}

func registerWorker(w http.ResponseWriter, r *http.Request) {
}

func upateWorker(w http.ResponseWriter, r *http.Request) {
}
