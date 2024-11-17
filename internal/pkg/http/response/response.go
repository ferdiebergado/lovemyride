package response

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type ValidationError struct {
	Field string
	Error string
}

type APIResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
	Data    any               `json:"data,omitempty"`
	Meta    map[string]any    `json:"meta,omitempty"`
}

func JSON[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func ServerError(w http.ResponseWriter, msg string, err error) {
	const errorText = "An error occurred."

	slog.Error(msg, "Error:", err)

	res := &APIResponse{
		Success: false,
		Message: errorText,
	}

	err = JSON(w, http.StatusInternalServerError, res)

	if err != nil {
		slog.Error(msg, "Error:", err)
		http.Error(w, errorText, http.StatusInternalServerError)
	}
}
