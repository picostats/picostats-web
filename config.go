package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	WebUrl        string `json:"web_url"`
	AppUrl        string `json:"app_url"`
	ListenAddr    string `json:"listen_addr"`
	Dev           bool   `json:"dev"`
	FSPrivateKey  string `json:"fs_private_key"`
	FSPrivateKey2 string `json:"fs_private_key2"`
	DBType        string `json:"db_type"`
	DBUrl         string `json:"db_url"`
	LogSQL        bool   `json:"log_sql"`
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
