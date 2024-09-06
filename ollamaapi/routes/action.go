package routes

import (
  "github.com/gin-gonic/gin"
  "1nsomnes/ollamaapi/webcalls"
  "net/http"
)

type Prompt struct {
  Prompt string `json:"prompt" binding:"required"`
  Actions map[string]string `json:"actions" binding:"required"`
}

func RegisterActionRoute(r *gin.Engine) {
  r.POST("/action", func(c *gin.Context) {
    var prompt Prompt

    if err := c.ShouldBindJSON(&prompt); err != nil {

      c.JSON(http.StatusBadRequest, gin.H{
        "error": err.Error(),
      })
      return
    }
    
    response := webcalls.CallGemma2b(prompt.Prompt)

    c.JSON(http.StatusOK, gin.H{
      "message":response,
    })

  })

}

