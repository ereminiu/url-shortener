package main

import (
	"log"

	"github.com/ereminiu/link-shorter/pkg/handlers"
	"github.com/ereminiu/link-shorter/pkg/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	// init database
	err := repository.InitDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalln(err)
	}

	// start server
	r := gin.Default()

	r.POST("/addlink", handlers.CreateLink)
	r.GET("/getlink", handlers.GetLink)

	r.POST("/addcustomlink", handlers.CreateCustomLink)
	r.GET("/getcustomlink", handlers.GetCustomLink)

	r.Run(":1024")
}
