package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"main.go/config"
	"main.go/models"
	"main.go/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBconfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})
	r := routes.Setuprouter()
	//running
	r.Run()
}
