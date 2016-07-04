package main

import (
    "os"
    "github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")
    router := gin.New()
    router.Use(gin.Logger())
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    router.Run(":" + port)
}
