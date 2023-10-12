package res

import (
	"encoding/json"
	"net/http"
)

type ValidationErrorResponse struct {
	Error            string              `json:"error"`
	ValidationErrors map[string][]string `json:"validationErrors"`
}

func NewValidationErrorResponse(validationErrors map[string][]string) ValidationErrorResponse {
	return ValidationErrorResponse{Error: "validation error", ValidationErrors: validationErrors}
}

func (r ValidationErrorResponse) Send(w http.ResponseWriter) {
	// TODO: handle error
	// TODO: set response encoding
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(r)
}
