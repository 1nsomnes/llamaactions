package main

import (
  "1nsomnes/ollamaapi/routes"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  
  routes.RegisterRoutes(r)

  r.Run()
}
