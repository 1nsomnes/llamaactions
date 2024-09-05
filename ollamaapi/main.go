package main

import (
  "fmt"
  "1nsomnes/ollamaapi/webcalls"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  r.Run()

  fmt.Println(webcalls.CallGemma2b("Respond with one word: \"yes\""))
}
