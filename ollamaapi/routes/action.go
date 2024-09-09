package routes

import (
  "github.com/gin-gonic/gin"
  "1nsomnes/ollamaapi/webcalls"
  "net/http"
  "fmt"
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
    
    promptMsg := "Please categorize a sentence, I will first provide a sentence that you must categorize and then a list of actions each with a description. Respond ONLY WITH THE ACTION NAME that the sentence corresponds to or the word \"uncertain\" if you don't know."
    promptMsg += "\nSentence: " + prompt.Prompt + "\n"

    for action, description := range prompt.Actions {
      promptMsg += fmt.Sprintf("\n\"%s\":\"%s\"", action, description) 
    }
    
    print(promptMsg)
    response := webcalls.CallGemma2b(promptMsg)

    c.JSON(http.StatusOK, gin.H{
      "message":response,
    })

  })

}

