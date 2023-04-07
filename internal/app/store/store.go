package store

type Store interface {
	UserAuthPattern() UserAuthPatternRepository
	Service() ServiceRepository
	Device() DeviceRepository
	User() UserRepository
	System() SystemRepository
}
