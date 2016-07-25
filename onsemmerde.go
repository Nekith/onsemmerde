package main

import (
    "bytes"
    "os"
    "fmt"
//    "log"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
//    "gopkg.in/mgo.v2/bson"
)

func Mongo() gin.HandlerFunc {
    return func(c *gin.Context) {
        session, exists := c.Get("mongo")
        if exists == false {
            var buffer bytes.Buffer
            buffer.WriteString("mongodb://")
            buffer.WriteString(os.Getenv("MONGODB_USER"))
            buffer.WriteString(":")
            buffer.WriteString(os.Getenv("MONGODB_PASSWORD"))
            buffer.WriteString("@")
            buffer.WriteString(os.Getenv("MONGODB_URI"))
            session, err := mgo.Dial(buffer.String())
            if err != nil {
                panic(err)
            }
            c.Set("mongo", session)
        }
        fmt.Println(session)
    }
}

func main() {
    port := os.Getenv("PORT")
    router := gin.Default()
    router.Use(Mongo())
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    router.Run(":" + port)
}
