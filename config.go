package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AppUrl     string `json:"app_url"`
	ListenAddr string `json:"listen_addr"`
	Dev        bool   `json:"dev"`
}

func (c *Config) load(configFile string) error {
	file, err := os.Open(configFile)

	if err != nil {
		log.Printf("[config.go] Error while opening config file: %v", err)
		return err
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&c)

	if err != nil {
		log.Printf("[config.go] Error while decoding JSON: %v", err)
		return err
	}

	return nil
}

func initConfig() *Config {
	conf := &Config{}
	conf.load("config.json")
	return conf
}
