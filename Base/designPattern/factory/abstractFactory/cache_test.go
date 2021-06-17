package abstractFactory

import (
	"testing"
	"fmt"
)

func TestRedisCacheFactory_Create(t *testing.T) {
	redisCacheFactory := &RedisCacheFactory{}
	redisCache := redisCacheFactory.Create()
	k := "K1"
	redisCache.Set(k, "V1")
	fmt.Println(redisCache.Get(k))
}

func TestMemCacheFactory_Create(t *testing.T) {
	memCacheFactory := &MemCacheFactory{}
	memCache := memCacheFactory.Create()
	k := "K1"
	memCache.Set(k, "V1")
	fmt.Println(memCache.Get(k))
}
