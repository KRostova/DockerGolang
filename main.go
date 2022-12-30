package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Str(ctx *gin.Context) {
	username, exists := os.LookupEnv("USER_NAME")

	if !exists {
		username = "Roman"
	}

	ctx.String(http.StatusOK, fmt.Sprintf("Hello %s!", username))
}

func Json(ctx *gin.Context) {
	username, exists := os.LookupEnv("USER_NAME")

	if !exists {
		username = "Roman"
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Hello!", "username": username})
}

func main() {
	portVariable := "SERVICE_PORT"
	port, exists := os.LookupEnv(portVariable)

	if !exists {
		fmt.Printf("[ERROR] Environment variable %s is not set.\n", portVariable)
		return
	}

	fmt.Printf("[INFO] Listening on port %s.\n", port)
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()

	e.GET("/string", Str)
	e.GET("/json", Json)

	e.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
