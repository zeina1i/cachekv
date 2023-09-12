package cachekv

var totalSize int
var evictionPool []string

func zmalloc(size int) []byte {
	if totalSize+size > 1000000000 {
		panic("out of memory")
	}

	totalSize += size

	return make([]byte, size)
}

func zfree(b []byte) {
	b = nil
}

func fillEvictionPool(c *Cache) {
	sampleSize := 100
	evictionPool = make([]string, sampleSize)

	for k, _ := range c.m {
		evictionPool = append(evictionPool, k)
	}
}

func evict() {

}

type Item struct {
	value   string
	lastUse int64
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
	zmalloc(len(value))
	c.m[key] = value
}

func (c *Cache) Del(key string) {
	zfree(c.m[key])
	delete(c.m, key)
}

func (c *Cache) Clear() {
	c.m = make(map[string][]byte)
}
