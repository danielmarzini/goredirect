package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/*", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, os.Getenv("REDIRECT_DESTINATION"))
	})
	r.Run()
}
