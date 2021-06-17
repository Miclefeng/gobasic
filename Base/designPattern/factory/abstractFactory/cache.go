package abstractFactory

type Cache interface {
	Set(k, v string)
	Get(k string) string
}

type RedisCache struct {
	data map[string]string
}

func NewRedisCache() *RedisCache {
	return &RedisCache{data: map[string]string{}}
}

func (r *RedisCache) Set(k, v string) {
	r.data[k] = v
}

func (r *RedisCache) Get(k string) string {
	var (
		v  string
		ok bool
	)

	if v, ok = r.data[k]; ok {
		return "Redis: " + v
	}
	return ""
}

type MemCache struct {
	data map[string]string
}

func NewMemCache() *MemCache {
	return &MemCache{data: map[string]string{}}
}

func (m *MemCache) Set(k, v string) {
	m.data[k] = v
}

func (m *MemCache) Get(k string) string {
	var (
		v  string
		ok bool
	)

	if v, ok = m.data[k]; ok {
		return "Mem: " + v
	}
	return ""
}

type CacheFactory interface {
	Create() Cache
}

type RedisCacheFactory struct{}

func (rc *RedisCacheFactory) Create() Cache {
	return NewRedisCache()
}

type MemCacheFactory struct{}

func (mf *MemCacheFactory) Create() Cache {
	return NewMemCache()
}
