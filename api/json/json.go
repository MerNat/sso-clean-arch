package json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GenericResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code"`
}

type Error struct {
	CodeStatus string `json:"code_status"`
	Message    string `json:"message,omitempty"`
	Err        string `json:"error,omitempty"`
}

const (
	ECONFLICT     = "conflict"            // action cannot be performed
	EINTERNAL     = "internal"            // internal error
	EINVALID      = "invalid"             // validation failed
	ENOTFOUND     = "not_found"           // entity does not exist
	EUNAUTHORIZED = "unauthorized access" // unauthorized
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
