package main

import (
  "fmt"
  "1nsomnes/ollamaapi/webcalls"
  "1nsomnes/ollamaapi/routes"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  
  routes.RegisterRoutes(r)

  r.Run()

  fmt.Println(webcalls.CallGemma2b("Respond with one word: \"yes\""))
}
