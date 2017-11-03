package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/YongMan/http4shell/auth"
	"github.com/YongMan/http4shell/config"
	"github.com/YongMan/http4shell/http"
	"github.com/gin-gonic/gin"
)

var (
	confFile string
	tokenGen bool
	username string
)

func init() {
	flag.StringVar(&confFile, "c", "", "config file path")
	flag.BoolVar(&tokenGen, "t", false, "generate token only, do not start server")
	flag.StringVar(&username, "u", "", "username for generate token")

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

	token := &auth.Token{Secret: c.Secret}
	if tokenGen {
		t, err := token.GenToken(username)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("User", username, "token is:")
		fmt.Println(t)
		return
	}
	r := gin.Default()

	authWrapper := auth.NewAuthWrapper(token)
	runHandler := http.NewHandlerRun(c)

	r.POST("/run", authWrapper.AuthWrapper(runHandler.RunHandler))

	// Listen and Server
	r.Run(c.Listen)
}
