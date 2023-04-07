package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
)

// handleServicesCreate ...
func (s *server) handleServicesCreate() http.HandlerFunc {
	type request struct {
		ServiceName string `json:"name"`
		ServiceSNI  string `json:"sni"`
	}

	return func(w http.ResponseWriter, req *http.Request) {
		r := &request{}

		if err := json.NewDecoder(req.Body).Decode(r); err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}

		service := &model.Service{
			ServiceName: r.ServiceName,
			ServiceSNI:  r.ServiceSNI,
		}

		if err := s.store.Service().Create(service); err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusCreated, service)

	}
}

// handleServicesDelete ...
func (s *server) handleServicesDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		idToDelete, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.Service().Delete(idToDelete); err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusNoContent, nil)

	}
}

// handleServicesFindAll ...
func (s *server) handleServicesFindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		services, err := s.store.Service().FindAll()
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(services)
	}
}

// handleServicesFindBySNI ...
func (s *server) handleServicesFindBySNI() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		sni, ok := vars["sni"]
		if !ok {
			s.error(w, req, http.StatusUnprocessableEntity, ErrUserNotFound)
			return
		}

		service, err := s.store.Service().FindBySNI(sni)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(service)
	}
}
