package utils

import (
	"encoding/json"
	"net/http"
)

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	NextOffset int         `json:"nextOffset,omitempty"`
}

func RespondPaginated(w http.ResponseWriter, data interface{}, total int64, nextOffset int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PaginatedResponse{
		Success:    true,
		Data:       data,
		Total:      total,
		NextOffset: nextOffset,
	})
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondSuccess(w http.ResponseWriter, data interface{}) {
	response := APIResponse{
		Success: true,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RespondError(w http.ResponseWriter, status int, errMsg string) {
	response := APIResponse{
		Success: false,
		Error:   errMsg,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
