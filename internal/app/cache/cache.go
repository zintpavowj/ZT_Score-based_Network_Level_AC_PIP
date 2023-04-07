package cache

import "github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"

type Cache struct {
	devices                   map[int]model.Device
	deviceUsedAuthPatterns    map[string][]string
	deviceTrustHistory        map[string][]float32
	deviceLocationIPHistory   map[string][]string
	deviceServiceUsageHistory map[string][]string
	deviceUserUsageHistory    map[string][]string

	services map[int]model.Service

	system model.System

	users                    map[int]model.User
	userUsedAuthPatterns     map[string][]string
	userTrustHistory         map[string][]float32
	userAccessRateHistory    map[string][]float32
	userInputBehaviorHistory map[string][]float32
	userServiceUsageHistory  map[string][]string
	userDeviceUsageHistory   map[string][]string
}

func New() *Cache {
	return &Cache{
		devices:                   make(map[int]model.Device),
		deviceUsedAuthPatterns:    make(map[string][]string),
		deviceTrustHistory:        make(map[string][]float32),
		deviceLocationIPHistory:   make(map[string][]string),
		deviceServiceUsageHistory: make(map[string][]string),
		deviceUserUsageHistory:    make(map[string][]string),

		services: make(map[int]model.Service),

		system: model.System{
			State:              "",
			PatchLevel:         "",
			ThreatLevel:        0,
			NetworkState:       "",
			NetworkThreatLevel: 0,
		},

		users:                    make(map[int]model.User),
		userUsedAuthPatterns:     make(map[string][]string),
		userTrustHistory:         make(map[string][]float32),
		userAccessRateHistory:    make(map[string][]float32),
		userInputBehaviorHistory: make(map[string][]float32),
		userServiceUsageHistory:  make(map[string][]string),
		userDeviceUsageHistory:   make(map[string][]string),
	}
}
