package main

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/kataras/iris.v6"
)

var app *iris.Framework

var conf *Config

var db *gorm.DB

func main() {
	// Loads and parses config.json file to struct
	conf = initConfig()

	// Initializes Iris web framework
	app = initIris()

	// Connects to the database
	db = initDB()

	// GET view handlers
	app.Get("/", pageView)
	app.Get("/pricing", pagePricingView)
	app.Get("/pricing/{userId}", pagePricingView)
	app.Get("/{page}", pageView)

	// POST view handlers
	app.Post("/paid", paidView)
	app.Post("/paid-cancel", paidCancelView)

	app.Listen(conf.ListenAddr)
}
