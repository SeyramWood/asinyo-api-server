package cache

import (
	"context"
	"errors"
	"log"
	"time"

	rcache "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"

	"github.com/SeyramWood/app/adapters/gateways"
)

var (
	KeyNotFoundErr = errors.New("key not found")
)

type cache struct {
	client *rcache.Cache
	ab     string
}

func New() gateways.CacheService {
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
	mycache := rcache.New(
		&rcache.Options{
			Redis:      rdb,
			LocalCache: rcache.NewTinyLFU(1000, time.Minute),
		},
	)
	return &cache{client: mycache, ab: "skljdhfklshd"}
}

func (c *cache) Set(key string, value any, ttl time.Duration) error {
	log.Println(key, c.ab)
	if err := c.client.Set(
		&rcache.Item{
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
func (c *cache) Get(key string, obj any) error {
	if err := c.client.Get(context.Background(), key, obj); err != nil {
		return KeyNotFoundErr
	}
	return nil
}
func (c *cache) Exist(key string) bool {
	return c.client.Exists(context.Background(), key)
}
func (c *cache) Delete(key string) error {
	return c.client.Delete(context.Background(), key)
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
