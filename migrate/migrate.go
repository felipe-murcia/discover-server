package main

import (
	"discoverco.co/server/configs"
	"discoverco.co/server/models"
	"discoverco.co/server/seeds"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
	seeds.CreateInitialFeatures()
}
