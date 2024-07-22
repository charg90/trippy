package main

import (
	"trippy/config"
	"trippy/db"
	"trippy/router"
)


func main() {
	config.LoadEnv()
	db.ConnectDb()
	r := router.SetupRouter()

	r.Run()

}