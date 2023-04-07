package cache

import (
	"errors"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

func (c *Cache) DeviceCreate(d *model.Device) {
	c.devices[d.ID] = *d
}

func (c *Cache) DeviceDelete(id int) {
	delete(c.devices, id)
}

func (c *Cache) DeviceFind(id int) (*model.Device, error) {
	value, ok := c.devices[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return &value, nil
}

func (c *Cache) DeviceFindByCN(cn string) (*model.Device, error) {
	for _, d := range c.devices {
		if d.DeviceCertCN == cn {
			return &d, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (c *Cache) DeviceUpdate(d *model.Device) {
	c.devices[d.ID] = *d
}

func (c *Cache) DeviceClear() {
	for _, d := range c.devices {
		delete(c.devices, d.ID)
	}
}

func (c *Cache) DeviceFindUsedAuthPatterns(cn string) ([]string, error) {
	value, ok := c.deviceUsedAuthPatterns[cn]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) DeviceSetUsedAuthPatterns(cn string, history []string) error {
	if c.deviceUsedAuthPatterns == nil {
		return errors.New("nil pointer assignment")
	}
	c.deviceUsedAuthPatterns[cn] = history
	return nil
}

func (c *Cache) DeviceFindTrustHistory(cn string) ([]float32, error) {
	value, ok := c.deviceTrustHistory[cn]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) DeviceSetTrustHistory(cn string, history []float32) error {
	if c.deviceTrustHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.deviceTrustHistory[cn] = history
	return nil
}

func (c *Cache) DeviceFindLocationIPHistory(cn string) ([]string, error) {
	value, ok := c.deviceLocationIPHistory[cn]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) DeviceSetLocationIPHistory(cn string, history []string) error {
	if c.deviceLocationIPHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.deviceLocationIPHistory[cn] = history
	return nil
}

func (c *Cache) DeviceFindServiceUsageHistory(cn string) ([]string, error) {
	value, ok := c.deviceServiceUsageHistory[cn]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) DeviceSetServiceUsageHistory(cn string, history []string) error {
	if c.deviceServiceUsageHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.deviceServiceUsageHistory[cn] = history
	return nil
}

func (c *Cache) DeviceFindUserUsageHistory(cn string) ([]string, error) {
	value, ok := c.deviceUserUsageHistory[cn]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}

func (c *Cache) DeviceSetUserUsageHistory(cn string, history []string) error {
	if c.deviceUserUsageHistory == nil {
		return errors.New("nil pointer assignment")
	}
	c.deviceUserUsageHistory[cn] = history
	return nil
}
