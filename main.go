package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	if os.Getenv("REDIRECT_DESTINATION") == "" {
		fmt.Println("please initialize env REDIRECT_DESTINATION")
		os.Exit(1)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, os.Getenv("REDIRECT_DESTINATION"))
	})
	r.Run()
}
