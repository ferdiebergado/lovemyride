package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FormData struct {
	Action string
	Method string
	Values any
}

func JSON[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
