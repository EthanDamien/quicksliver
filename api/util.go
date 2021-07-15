package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func RequireEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Envvar %s is required.", name)
	}

	return value
}

func sendErrorResponse(w http.ResponseWriter, httpStatus int, message string) {
	response := &response{
		Status: "failure",
		Message: "An error has occured",
		Error: message,
		Data: struct{}{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	writeJsonResponse(w, response)
}

func writeJsonResponse(w http.ResponseWriter, response interface{}) {
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Printf("cannot encode JSON response: %s", err.Error())
		sendInternalErrorResponse(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func sendInternalErrorResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, `{"status": 500, "message": "ERROR", "error": "Internal Service Error"}`, 500)
}
