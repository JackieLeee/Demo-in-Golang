package bigcache

import (
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
)

func TestBigCache(t *testing.T) {
	// 初始化缓存
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	// 写入缓存
	if err = cache.Set("key", []byte("value")); err != nil {
		t.Fatal(err)
	}

	// 读取缓存
	entry, err := cache.Get("key")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(entry))
}
