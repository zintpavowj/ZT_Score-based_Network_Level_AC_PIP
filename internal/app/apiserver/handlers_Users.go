package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
)

var (
	ErrUserNotFound error = errors.New("user not found")
)

// handleUsersCreate ...
func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Name           string    `json:"name"`
		Email          string    `json:"email"`
		LastAccessTime time.Time `json:"last_access_time"`
		Expected       float32   `json:"expected"`
		AccessTimeMin  string    `json:"access_time_min"`
		AccessTimeMax  string    `json:"access_time_max"`
	}

	return func(w http.ResponseWriter, req *http.Request) {
		r := &request{}

		if err := json.NewDecoder(req.Body).Decode(r); err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUsersCreate",
				"comment":  "unable to decode the request body",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}

		tmin, err := parseTime_HHMM(r.AccessTimeMin)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUsersCreate",
				"comment":  "unable to parse tmin",
			}).Error(err)

			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		tmax, err := parseTime_HHMM(r.AccessTimeMax)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUsersCreate",
				"comment":  "unable to parse tmax",
			}).Error(err)

			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		user := &model.User{
			Name:           r.Name,
			Email:          r.Email,
			LastAccessTime: r.LastAccessTime,
			Expected:       r.Expected,
			AccessTimeMin:  tmin,
			AccessTimeMax:  tmax,
		}

		if err := s.store.User().Create(user); err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUsersCreate",
				"comment":  "tmax",
			}).Error(err)

			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusCreated, user)

	}
}

// handleUsersDelete ...
func (s *server) handleUsersDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		idToDelete, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUsersDelete",
				"comment":  "unable to convert userID to int",
			}).Error(err)

			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.User().Delete(idToDelete); err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUsersDelete",
				"comment":  "unable to delete the user",
			}).Error(err)

			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, req, http.StatusNoContent, nil)

	}
}

// handleUsersFindByName ...
func (s *server) handleUsersFindByName() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		name, ok := vars["username"]
		if !ok {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUserUsedAuthPatterns",
			}).Error("no 'username' field in the request")

			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		user, err := s.store.User().FindByName(name)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "FindByName",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// handleUsersFindAll ...
func (s *server) handleUsersFindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		users, err := s.store.User().FindAll()
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUsersFindAll",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// handleUserUsedAuthPatterns ...
func (s *server) handleUserUsedAuthPatterns() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		userName, ok := vars["username"]
		if !ok {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUserUsedAuthPatterns",
			}).Error("no 'username' field in the request")

			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		userAuthPatternsStr, err := s.store.User().FindUsedAuthPatterns(userName)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "FindUsedAuthPatterns",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userAuthPatternsStr)
	}
}

// handleUserTrustHistory ...
func (s *server) handleUserTrustHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		userName, ok := vars["username"]
		if !ok {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUserTrustHistory",
			}).Error("no 'username' field in the request")

			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		trustHistory, err := s.store.User().FindTrustHistory(userName)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "FindTrustHistory",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(trustHistory)
	}
}

// handleUserTrustHistory ...
func (s *server) handleUserAccessRateHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		userName, ok := vars["username"]
		if !ok {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUserAccessRateHistory",
			}).Error("no 'username' field in the request")

			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		accessRateHistory, err := s.store.User().FindAccessRateHistory(userName)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "FindAccessRateHistory",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(accessRateHistory)
	}
}

// handleUserInputBehaviorHistory ...
func (s *server) handleUserInputBehaviorHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		userName, ok := vars["username"]
		if !ok {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUserInputBehaviorHistory",
			}).Error("no 'username' field in the request")

			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		inputBehaviorHistory, err := s.store.User().FindInputBehaviorHistory(userName)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "FindInputBehaviorHistory",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(inputBehaviorHistory)
	}
}

// handleUserServiceUsage ...
func (s *server) handleUserServiceUsage() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		userName, ok := vars["username"]
		if !ok {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUserServiceUsage",
			}).Error("no 'username' field in the request")

			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		serviceUsage, err := s.store.User().FindServiceUsageHistory(userName)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "FindServiceUsageHistory",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(serviceUsage)
	}
}

// handleUserDeviceUsage ...
func (s *server) handleUserDeviceUsage() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		userName, ok := vars["username"]
		if !ok {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "handleUserDeviceUsage",
			}).Error("no 'username' field in the request")

			s.error(w, req, http.StatusBadRequest, ErrUserNotFound)
			return
		}

		deviceUsage, err := s.store.User().FindDeviceUsageHistory(userName)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "apiserver",
				"function": "FindDeviceUsageHistory",
			}).Error(err)

			s.error(w, req, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deviceUsage)
	}
}

// parseTime_HHMM ...
func parseTime_HHMM(data string) (time.Time, error) {

	t := time.Time{}

	colonPosition := strings.Index(data, ":")
	if colonPosition == -1 {
		return t, errors.New("wrong time format")
	}

	hStr := data[:colonPosition]

	h, err := strconv.Atoi(hStr)
	if err != nil {
		return t, fmt.Errorf("wrong time value: %s", err.Error())
	}
	if (h < 0) || (h > 23) {
		return t, fmt.Errorf("wrong time value")
	}

	mStr := data[colonPosition+1:]

	m, err := strconv.Atoi(mStr)
	if err != nil {
		return t, fmt.Errorf("wrong time value: %s", err.Error())
	}
	if (h < 0) || (h > 60) {
		return t, fmt.Errorf("wrong time value")
	}

	t = time.Date(0, 0, 0, h, m, 0, 0, time.Local)
	return t, nil
}
