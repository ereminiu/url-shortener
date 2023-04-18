package main

import (
	"io/ioutil"
	"log"

	"github.com/ereminiu/url-shortener/pkg/handlers"
	"github.com/ereminiu/url-shortener/pkg/repository"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	DBName        string `yaml:"DBname"`
	SSLMode       string `yaml:"sslmode"`
	LocalHostPort string `yaml:"localhost_port"`
}

func main() {
	// read configs
	data, err := ioutil.ReadFile("../configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config.yaml")
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal configs")
	}

	// init database
	err = repository.InitDB(repository.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username,
		Password: cfg.Password,
		DBName:   cfg.DBName,
		SSLMode:  cfg.SSLMode,
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

	r.Run(cfg.LocalHostPort)
}
