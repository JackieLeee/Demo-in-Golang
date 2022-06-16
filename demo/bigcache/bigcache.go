package main

import (
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

func main() {
	// 初始化缓存
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Fatal(err)
	}

	// 写入缓存
	if err := cache.Set("key", []byte("value")); err != nil {
		log.Fatal(err)
	}

	// 读取缓存
	entry, err := cache.Get("key")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(entry))
}
