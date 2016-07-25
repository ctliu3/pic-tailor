package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Master       string  `yaml:"master_addr"`
	PingInterval int32   `yaml:"ping_interval"` // ms
	CPUUtil      float32 `yaml:"cpu_util"`      // float in (0, 1)
}

func loadConf(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	log.Printf("Load config %s succ.", filename)
	return &config, nil
}
