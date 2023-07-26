package app

import (
	"encoding/json"
	"net/http"

	"github.com/MidhunRajeevan/notification/model"
)

func BadRequest(w http.ResponseWriter, description string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	errorMessage := model.Error{Code: http.StatusBadRequest, Message: description}
	json.NewEncoder(w).Encode(errorMessage)

}
