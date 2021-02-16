// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package response provides functions to respond to an http request with
// information like Json objects or an error.
package response

import (
  "encoding/json"
  "net/http"
)

// Responds <w> with the error code <code> and a Json "message": "<msg>"
func RespondError(w http.ResponseWriter, code int, msg string){
  RespondJSON(w, code, map[string]string{"message": msg})
}

// Responds <w> with the code <code> and the Json representation of <payload>
func RespondJSON(w http.ResponseWriter, code int, payload interface{}){
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
