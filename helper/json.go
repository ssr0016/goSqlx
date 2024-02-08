package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(w http.ResponseWriter, r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(result); err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
}

func WriteResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}
