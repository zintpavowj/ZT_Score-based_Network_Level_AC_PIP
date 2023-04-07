package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
)

// handleUserAuthPatternsCreate ...
func (s *server) handleUserAuthPatternsCreate() http.HandlerFunc {
	type request struct {
		UserAuthPatternName string `json:"name"`
	}

	return func(w http.ResponseWriter, req *http.Request) {
		r := &request{}

		if err := json.NewDecoder(req.Body).Decode(r); err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}

		uap := &model.UserAuthPattern{
			UserAuthPatternName: r.UserAuthPatternName,
		}

		if err := s.store.UserAuthPattern().Create(uap); err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusCreated, uap)

	}
}

// handleUserAuthPatternsDelete ...
func (s *server) handleUserAuthPatternsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		idToDelete, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.UserAuthPattern().Delete(idToDelete); err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusNoContent, nil)

	}
}

// handleUserAuthPatternsFindAll ...
func (s *server) handleUserAuthPatternsFindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uaps, err := s.store.UserAuthPattern().FindAll()
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(uaps)
	}
}
