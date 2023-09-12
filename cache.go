package cachekv

var totalSize int
var limit int = 100

func addTotalSizeAndMayEvict(b []byte) {
	size := len(b)
	if totalSize+size > limit {
		panic("out of memory")
	}

	totalSize += size
}

func freeTotalSize(b []byte) {
	totalSize -= len(b)
}

type Cache struct {
	m map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string][]byte),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	v, ok := c.m[key]
	return v, ok
}

func (c *Cache) Set(key string, value []byte) {
	addTotalSizeAndMayEvict(value)
	c.m[key] = value
}

func (c *Cache) Del(key string) {
	freeTotalSize(c.m[key])
	delete(c.m, key)
}

func (c *Cache) Clear() {
	c.m = make(map[string][]byte)
}
