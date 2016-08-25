package main

import (
	"fmt"
	"html/template"
	// "io"
	"log"
	"net/http"
	// "time"

	"github.com/gorilla/mux"
)

const (
	STATIC_URL = "/static/"
	STATIC_DIR = "static/"
)

type WorkerState int

const (
	RUNNING WorkerState = iota
	STOP
	LOST
)

type Worker struct {
	WID          int32
	IP           string
	NumProcessed int32
	Statue       WorkerState
	LastUpdated  time
}

type Context struct {
	Title  string
	Static string
}

var (
	Workers       []Worker
	CPUUtilThresh float32
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_DIR))))

	router.HandleFunc("/", monitorHandler).Methods("GET")
	// Display the status of each worker on web.
	router.HandleFunc("/monitor", monitorHandler).Methods("GET")
	router.HandleFunc("/info", info).Methods("GET")
	router.HandleFunc("/workers", getWorkers).Methods("GET")

	router.HandleFunc("/health", registerWorker).Methods("POST")
	router.HandleFunc("/health", updateWorker).Methods("PUT")

	err := http.ListenAndServe(":8089", router)
	if err != nil {
		log.Fatal("Start server err: ", err)
	}
}

func render(w http.ResponseWriter, tpl string, context Context) {
	context.Static = STATIC_URL
	tpls := []string{"templates/base.html", fmt.Sprintf("templates/%s.html", tpl)}
	t, err := template.ParseFiles(tpls...)
	if err != nil {
		log.Fatal("Template parse err: ", err)
		return
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Fatal("Template execute err: ", err)
	}
}

func monitorHandler(w http.ResponseWriter, r *http.Request) {
	context := Context{Title: "monitor"}
	render(w, "monitor", context)
}

func info(w http.ResponseWriter, r *http.Request) {
	// Redirect reqeust to workers
}

func getWorkers(w http.ResponseWriter, r *http.Request) {
}

func registerWorker(w http.ResponseWriter, r *http.Request) {
}

func updateWorker(w http.ResponseWriter, r *http.Request) {
}
