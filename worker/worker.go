package main

import (
	"fmt"
	// "github.com/ctliu3/tailor/worker"
	"github.com/gorilla/mux"
	"image"
	_ "image/jpeg"
	_ "image/png"
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
	router.HandleFunc("/info", info).Methods("GET")

	go sendHearBeat()

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
