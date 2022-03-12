package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}

}

func Err(w http.ResponseWriter, statusCode int, err error) {
	Success(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}