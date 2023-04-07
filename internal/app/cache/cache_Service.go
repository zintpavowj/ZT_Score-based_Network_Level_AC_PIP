package cache

import (
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

func (c *Cache) ServiceCreate(s *model.Service) {
	c.services[s.ID] = *s
}

func (c *Cache) ServiceDelete(id int) {
	delete(c.services, id)
}

func (c *Cache) ServiceFind(id int) (*model.Service, error) {
	value, ok := c.services[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return &value, nil
}

func (c *Cache) ServiceFindBySNI(sni string) (*model.Service, error) {
	for _, s := range c.services {
		if s.ServiceSNI == sni {
			return &s, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (c *Cache) ServiceUpdate(s *model.Service) {
	c.services[s.ID] = *s
}

func (c *Cache) ServiceClear() {
	for _, s := range c.services {
		delete(c.services, s.ID)
	}
}
