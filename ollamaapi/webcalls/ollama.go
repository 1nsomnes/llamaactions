package webcalls

import (
  "net/http"
  "bytes"
  "encoding/json"
  "io/ioutil"
)

type Payload struct {
  Model string `json:"model"`
  Prompt string `json:"prompt"`
  Stream bool `json:"stream"`
}

type ResponseData struct {
//  Model string `json:"model"`
//  CreatedAt string `json:"created_ad"`
  Response string `json:"response"`
//  Done bool `json:"done:`
//  DoneReason string `json:"done_reason"`
//  Context []int `json:"context"`
//  TotalDuration int `json:"total_duration"`
//  LoadDuration  int `json:"load_duration"`
//  PromptEvalCount  int `json:"prompt_eval_count"`
//  PromptEvalDuration  int `json:"prompt_eval_duration"`
//  EvalCount  int `json:"eval_count"`
//  EvalDuration  int `json:"eval_duration"`
}

func CallGemma2b(msg string) string {
  url := "http://localhost:11434/api/generate" 
  
  data := Payload{
    Model: "gemma2:2b",
    Prompt: msg,
    Stream: false,
  }

  jsonData, _ := json.Marshal(data)

  req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, _ := client.Do(req)

  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return "Error: received non-OK http status: " + resp.Status
  }
  
  body, _ := ioutil.ReadAll(resp.Body)

  var responseData ResponseData
  err := json.Unmarshal(body, &responseData)
  if err != nil {
    return "Error parsing JSON response"
  }

  return responseData.Response
}
