package main

import (
	// "github.com/ctliu3/tailor/worker"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const (
	CONFIG_FILE = "conf.yaml"
)

var (
	config *Config
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		s := <-sigs
		log.Printf("Receive sigal %v, bye:-(", s)
		os.Exit(0)
	}()

	initialize()

	router := mux.NewRouter()
	router.HandleFunc("/meta", apiMeta).Methods("GET")
	router.HandleFunc("/image", apiImage).Methods("GET")

	go sendHearBeat()

	err := http.ListenAndServe(":"+config.WorkerPort, router)
	if err != nil {
		log.Fatal("Start server err: ", err)
	}
}

func initialize() {
	filename, err := filepath.Abs(CONFIG_FILE)
	if err != nil {
		panic(err)
	}
	config, err = loadConf(filename)
	if err != nil {
		panic(err)
	}
	log.Printf("config: %v", config)
}

func sendHearBeat() {
	var retry int32 = 0

	for {
		cpuInfo, err := getCPUInfo()
		if err != nil {
			log.Printf("%v", err)
			if retry == config.MaxRetry {
				break
			}
			retry += 1
		}
		log.Printf("cpu info %v", cpuInfo)
		time.Sleep(time.Second * time.Duration(rand.Int31n(config.PingInterval)))
	}

	log.Fatal("Reach max retry times, shut down.")
}

func monitorCPU() {
}
