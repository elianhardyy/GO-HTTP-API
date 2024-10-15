package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, code int, message string){
	JSONResponse(w,code,map[string]string{"error":message})
}
func JSONResponse(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(response)
}
