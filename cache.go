package cache

import "time"

type Cache struct {
	store map[string]string
	deadlines map[string]time.Time
}

func NewCache() Cache {
	cache := Cache{}
	cache.store = make(map[string]string)
	cache.deadlines = make(map[string]time.Time)
	return cache
}

func (c *Cache) Clear() {
	now := time.Now()
	for key, _ := range c.deadlines {
		if now.After(c.deadlines[key]) {
			delete(c.store, key)
		}
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.Clear()
	result, ok := c.store[key]
	return result, ok
}

func (c *Cache) Put(key, value string) {
	c.store[key] = value
}

func (c *Cache) Keys() []string {
	c.Clear()
	keys := make([]string, len(c.store))
	for key, _ := range c.store {
		keys[0] = key
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.store[key] = value
	c.deadlines[key] = deadline
}
