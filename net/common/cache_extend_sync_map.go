package common

type cacheExtend struct {
	pool Map
}

func NewCacheExtend() *cacheExtend {
	return &cacheExtend{}
}

func (c *cacheExtend) Get(domain string) (interface{}, bool) {
	return c.pool.Load(domain)
}

func (c *cacheExtend) Add(domain string, mark interface{}) {
	if mark == nil {
		return
	}
	c.pool.Store(domain, mark)

	if c.pool.Length() < 800 {
		return
	}

	c.pool.Range(func(key, value interface{}) bool {
		c.pool.Delete(key)
		if c.pool.Length() <= 700 {
			return false
		}
		return true
	})
}