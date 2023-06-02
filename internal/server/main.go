package main

import (
	"log"
	"vediomeeting/internal/models"
	"vediomeeting/internal/server/router"
)

func main() {
	models.NewDB()
	e := router.Router()
	err := e.Run()
	if err != nil {
		log.Fatalln("run err: ", err)
		return
	}
}
