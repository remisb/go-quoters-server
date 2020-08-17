package web

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

const (
	headerContentType   = "Content-Type"
	mimeApplicationJSON = "application/json"
)

func UrlParamInt(r *http.Request, key string, def int) int {
	val := chi.URLParam(r, key)
	if val != "" {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return def
		}
		return intVal
	}
	return def
}

// RespondError create json error response and outputs passed error into response body.
func RespondError(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	Respond(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...)},
	})
}

// Respond create json response and outputs json representation of the passed data into response body.
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set(headerContentType, mimeApplicationJSON)
	w.WriteHeader(status)
	if data != nil {
		EncodeBody(w, r, data)
	}
}

// EncodeBody encodes passed date to json format and writes it into Response body.
func EncodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
