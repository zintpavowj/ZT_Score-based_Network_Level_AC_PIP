package apiserver

// configureRouter() ...
func (s *server) configureRouter() {

	s.router.HandleFunc("/authpatterns/users", s.handleUserAuthPatternsFindAll()).Methods("GET")
	s.router.HandleFunc("/authpatterns/users", s.handleUserAuthPatternsCreate()).Methods("POST")
	s.router.HandleFunc("/authpatterns/users/{id:[0-9]+}", s.handleUserAuthPatternsDelete()).Methods("DELETE")

	s.router.HandleFunc("/services", s.handleServicesFindAll()).Methods("GET")
	s.router.HandleFunc("/services", s.handleServicesCreate()).Methods("POST")
	s.router.HandleFunc("/services/{id:[0-9]+}", s.handleServicesDelete()).Methods("DELETE")
	s.router.HandleFunc("/services/{sni}", s.handleServicesFindBySNI()).Methods("GET")

	s.router.HandleFunc("/devices", s.handleDevicesFindAll()).Methods("GET")
	s.router.HandleFunc("/devices", s.handleDevicesCreate()).Methods("POST")
	s.router.HandleFunc("/devices/{id:[0-9]+}", s.handleDevicesDelete()).Methods("DELETE")
	s.router.HandleFunc("/devices/{cert_cn}", s.handleDevicesFindByName()).Methods("GET")
	s.router.HandleFunc("/devices/{cert_cn}/authpatterns", s.handleDeviceUsedAuthPatterns()).Methods("GET")
	s.router.HandleFunc("/devices/{cert_cn}/trusthistory", s.handleDeviceTrustHistory()).Methods("GET")
	s.router.HandleFunc("/devices/{cert_cn}/iphistory", s.handleDeviceLocationIPHistory()).Methods("GET")
	s.router.HandleFunc("/devices/{cert_cn}/serviceusage", s.handleDeviceServiceUsage()).Methods("GET")
	s.router.HandleFunc("/devices/{cert_cn}/userusage", s.handleDeviceUserUsage()).Methods("GET")

	s.router.HandleFunc("/users", s.handleUsersFindAll()).Methods("GET")
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/users/{id:[0-9]+}", s.handleUsersDelete()).Methods("DELETE")
	s.router.HandleFunc("/users/{username}", s.handleUsersFindByName()).Methods("GET")
	s.router.HandleFunc("/users/{username}/authpatterns", s.handleUserUsedAuthPatterns()).Methods("GET")
	s.router.HandleFunc("/users/{username}/trusthistory", s.handleUserTrustHistory()).Methods("GET")
	s.router.HandleFunc("/users/{username}/accessratehistory", s.handleUserAccessRateHistory()).Methods("GET")
	s.router.HandleFunc("/users/{username}/inputbehavior", s.handleUserInputBehaviorHistory()).Methods("GET")
	s.router.HandleFunc("/users/{username}/serviceusage", s.handleUserServiceUsage()).Methods("GET")
	s.router.HandleFunc("/users/{username}/deviceusage", s.handleUserDeviceUsage()).Methods("GET")

	s.router.HandleFunc("/system", s.handleGetSystemState()).Methods("GET")
}
