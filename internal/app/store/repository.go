package store

import "github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"

// UserAuthPatternRepository ...
type UserAuthPatternRepository interface {
	Create(*model.UserAuthPattern) error
	Delete(int) error
	Find(int) (*model.UserAuthPattern, error)
	FindAll() ([]model.UserAuthPattern, error)
	Clear() error
}

// ServiceRepository ...
type ServiceRepository interface {
	Create(*model.Service) error
	Delete(int) error
	Find(int) (*model.Service, error)
	FindBySNI(string) (*model.Service, error)
	FindAll() ([]model.Service, error)
	Clear() error
}

// DeviceRepository ...
type DeviceRepository interface {
	Create(*model.Device) error
	Delete(int) error
	Find(int) (*model.Device, error)
	FindByCN(string) (*model.Device, error)
	FindAll() ([]model.Device, error)
	FindUsedAuthPatterns(string) ([]string, error)
	FindTrustHistory(string) ([]float32, error)
	FindLocationIPHistory(string) ([]string, error)
	FindServiceUsageHistory(string) ([]string, error)
	FindUserUsageHistory(string) ([]string, error)
	Clear() error
}

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Delete(int) error
	Find(int) (*model.User, error)
	FindByName(string) (*model.User, error)
	FindAll() ([]model.User, error)
	FindUsedAuthPatterns(string) ([]string, error)
	FindTrustHistory(string) ([]float32, error)
	FindAccessRateHistory(string) ([]float32, error)
	FindInputBehaviorHistory(string) ([]float32, error)
	FindServiceUsageHistory(string) ([]string, error)
	FindDeviceUsageHistory(string) ([]string, error)
	Clear() error
}

// SystemRepository ...
type SystemRepository interface {
	Get() (*model.System, error)
}
