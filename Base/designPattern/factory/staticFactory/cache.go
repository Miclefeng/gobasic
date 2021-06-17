package staticFactory

import "errors"

type Cache interface {
	Set(k, v string)
	Get(k string) string
}

type cacheType int

const (
	redisType cacheType = iota
	memType
)

type RedisCache struct {
	data map[string]string
}

func (r *RedisCache) Set(k, v string) {
	r.data[k] = v
}

func (r *RedisCache) Get(k string) string {
	var (
		v  string
		ok bool
	)
	if v, ok = r.data[k]; !ok {
		return ""
	}
	return "Redis: " + v
}

type MemCache struct {
	data map[string]string
}

func (m *MemCache) Set(k, v string) {
	m.data[k] = v
}

func (m *MemCache) Get(k string) string {
	var (
		v  string
		ok bool
	)
	if v, ok = m.data[k]; !ok {
		return ""
	}
	return "Mem: " + v
}

type CacheFactory struct{}

func (f *CacheFactory) Create(cacheType cacheType) (Cache, error) {
	if cacheType == redisType {
		return &RedisCache{data: map[string]string{}}, nil
	}

	if cacheType == memType {
		return &MemCache{data: map[string]string{}}, nil
	}

	return nil, errors.New("error cache type.")
}
