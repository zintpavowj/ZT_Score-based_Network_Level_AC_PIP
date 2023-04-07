package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
)

var (
	ErrDeviceNotFound error = errors.New("device not found")
)

// handleDevicesCreate ...
func (s *server) handleDevicesCreate() http.HandlerFunc {
	type request struct {
		DeviceName   string `json:"name"`
		DeviceCertCN string `json:"cert_cn"`
	}

	return func(w http.ResponseWriter, req *http.Request) {
		r := &request{}

		if err := json.NewDecoder(req.Body).Decode(r); err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}

		device := &model.Device{
			DeviceName:   r.DeviceName,
			DeviceCertCN: r.DeviceCertCN,
		}

		if err := s.store.Device().Create(device); err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusCreated, device)

	}
}

// handleDevicesDelete ...
func (s *server) handleDevicesDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		idToDelete, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.Device().Delete(idToDelete); err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusNoContent, nil)

	}
}

// handleDevicesFindAll ...
func (s *server) handleDevicesFindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		devices, err := s.store.Device().FindAll()
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(devices)
	}
}

// handleDevicesFindByName ...
func (s *server) handleDevicesFindByName() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		cert_cn, ok := vars["cert_cn"]
		if !ok {
			s.error(w, req, http.StatusUnprocessableEntity, ErrDeviceNotFound)
			return
		}

		device, err := s.store.Device().FindByCN(cert_cn)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(device)
	}
}

// handleDeviceUsedAuthPatterns ...
func (s *server) handleDeviceUsedAuthPatterns() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		deviceCN, ok := vars["cert_cn"]
		if !ok {
			s.error(w, req, http.StatusBadRequest, ErrDeviceNotFound)
			return
		}

		userAuthPatternsStr, err := s.store.Device().FindUsedAuthPatterns(deviceCN)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userAuthPatternsStr)
	}
}

// handleDeviceTrustHistory ...
func (s *server) handleDeviceTrustHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		cn, ok := vars["cert_cn"]
		if !ok {
			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		trustHistory, err := s.store.Device().FindTrustHistory(cn)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(trustHistory)
	}
}

// handleDeviceLocationIPHistory ...
func (s *server) handleDeviceLocationIPHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		cn, ok := vars["cert_cn"]
		if !ok {
			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		locationIPHistory, err := s.store.Device().FindLocationIPHistory(cn)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(locationIPHistory)
	}
}

// handleDeviceServiceUsage ...
func (s *server) handleDeviceServiceUsage() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		userName, ok := vars["cert_cn"]
		if !ok {
			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		serviceUsage, err := s.store.Device().FindServiceUsageHistory(userName)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(serviceUsage)
	}
}

// handleDeviceUserUsage ...
func (s *server) handleDeviceUserUsage() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		cn, ok := vars["cert_cn"]
		if !ok {
			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		userUsage, err := s.store.Device().FindUserUsageHistory(cn)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userUsage)
	}
}
