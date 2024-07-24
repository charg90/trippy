package main

import (
	"trippy/config"
	"trippy/db"
	"trippy/db/models"
)

func init () {
	config.LoadEnv()
	db.ConnectDb()
}

func main()  {
	db.DB.AutoMigrate(&models.User{})

}