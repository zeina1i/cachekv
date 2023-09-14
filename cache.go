package cachekv

import "time"

var totalSize int
var limit int = 100

type Item struct {
	value          []byte
	lastAccessTime int64
}
type Cache struct {
	m map[string]Item
	e *eviction
}

func NewCache() *Cache {
	m := make(map[string]Item)
	e := newEvictionPool()
	e.fillPoolPeriodically(m)
	return &Cache{
		m: m,
		e: newEvictionPool(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	v, ok := c.m[key]
	v.lastAccessTime = time.Now().Unix()
	return v.value, ok
}

func (c *Cache) Set(key string, value []byte) {
	item := Item{
		value:          value,
		lastAccessTime: 0,
	}
	c.addTotalSizeAndMayEvict(item.value)
	c.m[key] = item
}

func (c *Cache) Del(key string) {
	c.freeTotalSize(c.m[key].value)
	delete(c.m, key)
}

func (c *Cache) Clear() {
	c.m = make(map[string]Item)
}

func (c *Cache) addTotalSizeAndMayEvict(b []byte) {
	size := len(b) + 8
	for totalSize+size > limit {
		c.e.evict()
	}

	totalSize += size
}

func (c *Cache) freeTotalSize(b []byte) {
	totalSize -= len(b) - 8
}
