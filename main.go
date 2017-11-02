package main

import (
	"github.com/YongMan/http4shell/auth"
	"github.com/YongMan/http4shell/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/run", auth.AuthWrapper(http.RunHandler))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
