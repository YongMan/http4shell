package main

import (
	"flag"
	"log"

	"github.com/YongMan/http4shell/auth"
	"github.com/YongMan/http4shell/config"
	"github.com/YongMan/http4shell/http"
	"github.com/gin-gonic/gin"
)

var (
	confFile string
)

func init() {
	flag.StringVar(&confFile, "c", "", "config file path")
}

func main() {
	flag.Parse()

	if confFile == "" {
		log.Fatal("config file should be assigned")
	}

	c, err := config.LoadConfig(confFile)
	if err != nil {
		log.Fatal("config file format invalid")
	}

	r := gin.Default()

	authWrapper := auth.NewAuthWrapper()
	runHandler := http.NewHandlerRun(c)

	r.POST("/run", authWrapper.AuthWrapper(runHandler.RunHandler))

	// Listen and Server
	r.Run(c.Listen)
}
