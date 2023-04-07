package apiserver

import (
	"net/http"
)

// handleGetSystemState ...
func (s *server) handleGetSystemState() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		system, err := s.store.System().Get()
		if err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusOK, system)

	}
}
