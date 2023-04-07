package cache

import (
	"errors"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

func (c *Cache) UserCreate(u *model.User) {
	c.users[u.ID] = *u
}

func (c *Cache) UserDelete(id int) {
	delete(c.users, id)
}

func (c *Cache) UserFind(id int) (*model.User, error) {
	value, ok := c.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return &value, nil
}

func (c *Cache) UserFindByName(name string) (*model.User, error) {
	for _, u := range c.users {
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (c *Cache) UserUpdate(u *model.User) {
	c.users[u.ID] = *u
}

func (c *Cache) UserClear() {
	for _, u := range c.users {
		delete(c.users, u.ID)
	}
}

func (c *Cache) UserFindUsedAuthPatterns(name string) ([]string, error) {
	value, ok := c.userUsedAuthPatterns[name]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) UserSetUsedAuthPatterns(name string, history []string) error {
	if c.userUsedAuthPatterns == nil {
		return errors.New("nil pointer assignment")
	}
	c.userUsedAuthPatterns[name] = history
	return nil
}

func (c *Cache) UserFindTrustHistory(name string) ([]float32, error) {
	value, ok := c.userTrustHistory[name]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) UserSetTrustHistory(name string, history []float32) error {
	if c.userTrustHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.userTrustHistory[name] = history
	return nil
}

func (c *Cache) UserFindAccessRateHistory(name string) ([]float32, error) {
	value, ok := c.userAccessRateHistory[name]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) UserSetAccessRateHistory(name string, history []float32) error {
	if c.userAccessRateHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.userAccessRateHistory[name] = history
	return nil
}

func (c *Cache) UserFindInputBehaviorHistory(name string) ([]float32, error) {
	value, ok := c.userInputBehaviorHistory[name]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) UserSetInputBehaviorHistory(name string, history []float32) error {
	if c.userInputBehaviorHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.userInputBehaviorHistory[name] = history
	return nil
}

func (c *Cache) UserFindServiceUsageHistory(name string) ([]string, error) {
	value, ok := c.userServiceUsageHistory[name]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) UserSetServiceUsageHistory(name string, history []string) error {
	if c.userServiceUsageHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.userServiceUsageHistory[name] = history
	return nil
}

func (c *Cache) UserFindDeviceUsageHistory(name string) ([]string, error) {
	value, ok := c.userDeviceUsageHistory[name]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) UserSetDeviceUsageHistory(name string, history []string) error {
	if c.userDeviceUsageHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.userDeviceUsageHistory[name] = history
	return nil
}
