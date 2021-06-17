package staticFactory

import (
	"testing"
	"fmt"
)

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}
	redis, err := cacheFactory.Create(redisType)
	if err != nil {
		t.Error(err)
	}

	k := "K1"
	redis.Set(k, "V1")
	fmt.Println(redis.Get(k))

	mem, err := cacheFactory.Create(memType)
	if err != nil {
		t.Error(err)
	}

	mem.Set(k, "V1")
	fmt.Println(mem.Get(k))
}
