package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		ctx.Writer.Header().Set("Location", os.Getenv("REDIRECT_DESTINATION"))
		ctx.Abort(301)
	})
	r.Run()
}
