package app_cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type AppCache struct {
	client *cache.Cache
}

func New() *AppCache {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	)
	_, e := rdb.Ping(context.Background()).Result()
	if e != nil {
		log.Println(e)
	}
	mycache := cache.New(
		&cache.Options{
			Redis:      rdb,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		},
	)
	return &AppCache{client: mycache}
}

func (ac *AppCache) Set(key string, value any, ttl time.Duration) error {
	if err := ac.client.Set(
		&cache.Item{
			Ctx:   context.Background(),
			Key:   key,
			Value: value,
			TTL:   ttl,
		},
	); err != nil {
		return err
	}
	return nil
}
func (ac *AppCache) Get(key string, obj any) (any, error) {
	var object = obj
	if err := ac.client.Get(context.Background(), key, object); err != nil {
		return nil, err
	}
	return object, nil
}
func (ac *AppCache) Exist(key string) bool {
	return ac.client.Exists(context.Background(), key)
}
func (ac *AppCache) Delete(key string) error {
	return ac.client.Delete(context.Background(), key)
}

// CleanUp periodically removes expired entries from the cache.
// func (ac *AppCache) CleanUp() {
// 	for {
// 		time.Sleep(1 * time.Minute)
// 		ac.cache.Range(
// 			func(key, entry any) bool {
// 				cacheEntry := entry.(cacheEntry)
// 				if time.Now().UnixNano() > cacheEntry.expiration {
// 					ac.cache.Delete(key)
// 				}
// 				return true
// 			},
// 		)
// 	}
// }
