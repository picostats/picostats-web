package main

import (
	"gopkg.in/kataras/iris.v6"
)

var app *iris.Framework

var conf *Config

func main() {
	// Loads and parses config.json file to struct
	conf = initConfig()

	// Initializes Iris web framework
	app = initIris()

	app.Get("/", pageView)
	app.Get("/{page}", pageView)

	app.Listen(conf.ListenAddr)
}
