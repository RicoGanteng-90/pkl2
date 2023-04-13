package exception

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Errors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendErrorResponse(w http.ResponseWriter, errorCode int, errorMessage string) {
	response := Errors{
		Code:    errorCode,
		Message: errorMessage,
	}
	SendJsonResponse(w, errorCode, response)
}
func SendJsonResponse(w http.ResponseWriter, errorCode int, body interface{}) {
	response, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	w.Write(response)
}

func BearerUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	SendErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("Token is invalid"))
}
