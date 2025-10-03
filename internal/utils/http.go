package utils

import (
	"encoding/json"
	"github.com/kstsm/wb-shortener/internal/models"
	"net/http"
	"strings"
)

func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		return
	}

	resp := map[string]interface{}{
		"result": data,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, `{"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := models.Error{Error: message}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, `{"error":"failed to encode error response", "message": '}`, http.StatusInternalServerError)
	}
}

func GetClientIP(r *http.Request) string {
	headers := []string{"X-Forwarded-For", "X-Real-IP", "X-Forwarded", "Forwarded-For", "Forwarded"}
	for _, h := range headers {
		if ip := r.Header.Get(h); ip != "" {
			return strings.TrimSpace(strings.Split(ip, ",")[0])
		}
	}

	ip := r.RemoteAddr
	if colonIndex := strings.LastIndex(ip, ":"); colonIndex != -1 {
		ip = ip[:colonIndex]
	}
	return ip
}
