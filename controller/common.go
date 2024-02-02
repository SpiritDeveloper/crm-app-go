package controller

import (
	. "crm-app-go/docs"
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithArrayError(w http.ResponseWriter, code int, message []string) {
	response := &GenericError{
		Success: true,
		Message: "Oops, there seems to be a problem with your request",
		Code: code,
    }
	response.Errors.Error = message
	respondWithJSON(w, code, response)
}


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
