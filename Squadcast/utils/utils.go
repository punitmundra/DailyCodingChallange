package utils

import (
	"encoding/json"
	"net/http"
)

//Message return the maps
func Message(status bool, message string, code int) map[string]interface{} {
	return map[string]interface{}{"code": code, "status": status, "message": message}
}

// Respond return encode data
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
