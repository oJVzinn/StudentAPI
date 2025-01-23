package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"studentapi/controller"
	"studentapi/database"
)

func main() {
	database.Init()
	instance := gin.Default()
	controller.Setup(instance)
	err := instance.Run(":8080")
	log.Fatal(err)
}
