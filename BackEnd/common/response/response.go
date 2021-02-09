package response

import (
  "encoding/json"
  "net/http"
)

func RespondError(w http.ResponseWriter, code int, msg string){
  RespondJSON(w, code, map[string]string{"message": msg})
}

func RespondJSON(w http.ResponseWriter, code int, payload interface{}){
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
