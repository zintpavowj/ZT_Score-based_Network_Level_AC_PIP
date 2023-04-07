package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/cache"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

// Store ...
type Store struct {
	db                        *sql.DB
	c                         *cache.Cache
	useCache                  bool
	userAuthPatternRepository *UserAuthPatternRepository
	serviceRepository         *ServiceRepository
	deviceRepository          *DeviceRepository
	userRepository            *UserRepository
	systemRepository          *SystemRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db:       db,
		c:        nil,
		useCache: false,
	}
}

// The function initializes a pointer to the internal DB cache storage
func (s *Store) SetCache(c *cache.Cache) {
	s.c = c
}

// The function enables using of the DB cache
func (s *Store) EnableCache() {
	s.useCache = true
}

// The function disables using of the DB cache
func (s *Store) DisableCache() {
	s.useCache = false
}

// UserAuthPattern ...
func (s *Store) UserAuthPattern() store.UserAuthPatternRepository {
	if s.userAuthPatternRepository != nil {
		return s.userAuthPatternRepository
	}

	s.userAuthPatternRepository = &UserAuthPatternRepository{
		store: s,
	}

	return s.userAuthPatternRepository
}

// Service ...
func (s *Store) Service() store.ServiceRepository {
	if s.serviceRepository != nil {
		return s.serviceRepository
	}

	s.serviceRepository = &ServiceRepository{
		store: s,
	}

	return s.serviceRepository
}

// Device ...
func (s *Store) Device() store.DeviceRepository {
	if s.deviceRepository != nil {
		return s.deviceRepository
	}

	s.deviceRepository = &DeviceRepository{
		store: s,
	}

	return s.deviceRepository
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// System ...
func (s *Store) System() store.SystemRepository {
	if s.systemRepository != nil {
		return s.systemRepository
	}

	s.systemRepository = &SystemRepository{
		store: s,
	}

	return s.systemRepository
}
