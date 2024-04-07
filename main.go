package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {

    r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

    r.Run(":80")
}
