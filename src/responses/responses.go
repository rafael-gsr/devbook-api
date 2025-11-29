// Package responses contains all the api responses handlers
package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns the api response as a json object
func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)

	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}

// Error is the function that handlers error responses of the api
func Error(w http.ResponseWriter, statusCode int, error error) {
	errorStruct := struct {
		Error string `json:"error"`
	}{Error: error.Error()}

	JSON(w, statusCode, errorStruct)
}
