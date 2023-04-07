package cache

import (
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
)

func (c *Cache) SystemGet() *model.System {
	return &c.system
}

func (c *Cache) SystemSet(s *model.System) {
	c.system = *s
}
