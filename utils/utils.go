package utils

import (
	"encoding/json"
	"net/http"
)

func JsonifyHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func Wrap(w http.ResponseWriter, v map[string]interface{}) {
	JsonifyHeader(w)
	_ = json.NewEncoder(w).Encode(v)
}

func RespWrap(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"msg": msg})
}

func ErrWrap(w http.ResponseWriter, v string) {
	JsonifyHeader(w)
	w.WriteHeader(500)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": v})
}
