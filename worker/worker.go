package main

import (
	"fmt"
	// "github.com/ctliu3/tailor/worker"
	"github.com/gorilla/mux"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"time"
)

const (
	CONFIG_FILE = "conf.yaml"
)

var (
	config *Config
)

func main() {
	initialize()

	router := mux.NewRouter()
	router.HandleFunc("/info", info).Methods("GET")

	go connectMaster()
	err := http.ListenAndServe(":8090", router)
	if err != nil {
		log.Fatal("Start server err: ", err)
	}
}

func initialize() {
	filename, err := filepath.Abs(CONFIG_FILE)
	if err != nil {
		panic(err)
	}
	config, err := loadConf(filename)
	if err != nil {
		panic(err)
	}
	log.Printf("config: %v", config)
}

func connectMaster() {
	cpuInfo, err := getCPUInfo()
	if err != nil {
	}
	fmt.Printf("%v", cpuInfo)
	fmt.Printf("%v", reflect.TypeOf(cpuInfo))
}

func monitorCPU() {
}

func info(w http.ResponseWriter, r *http.Request) {
	image_data := r.Body
	if image_data == nil {
		fmt.Fprint(w, "err: body can not be empty")
		return
	}
	start := time.Now()
	m, _, err := image.Decode(image_data)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Fprint(w, "err: image can not be decoded")
		return
	}
	log.Printf("size: %v, elapsed: %v", m.Bounds(), elapsed)

	w.WriteHeader(http.StatusOK)
}
