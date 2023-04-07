package apiserver

import (
	"encoding/json"
	"net/http"
)

func (s *server) error(w http.ResponseWriter, req *http.Request, code int, err error) {
	s.respond(w, req, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, req *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
