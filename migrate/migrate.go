package main

import (
	"go-backend/configs"
	"go-backend/models"
)

func init() {
	configs.ConnectToDb()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}
