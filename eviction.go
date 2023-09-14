package cachekv

import "time"

var evictionPoolSize int = 100

type eviction struct {
	ticker                *time.Ticker
	len                   int
	maximumLastAccessTime int64
	pool                  []string
	done                  chan bool
}

func newEvictionPool() *eviction {
	return &eviction{
		pool: make([]string, 0, evictionPoolSize),
	}
}

func (e *eviction) fillPoolPeriodically(m map[string]Item) (*time.Ticker, chan bool) {
	go func() {
		for {
			select {
			case <-e.done:
				return
			case _ = <-e.ticker.C:
				for key, value := range m {
					if e.len >= evictionPoolSize {
						break
					}
					e.mayPush(key, value.lastAccessTime)
				}
			}
		}
	}()

	return e.ticker, e.done
}

func (e *eviction) mayPush(key string, lastAccessTime int64) {
	if lastAccessTime > e.maximumLastAccessTime {
		e.maximumLastAccessTime = lastAccessTime
		e.pool = append(e.pool, key)
		e.len++
	}
}

func (e *eviction) evict() string {
	key := e.pool[0]
	e.pool = e.pool[1:]
	e.len--

	return key
}
