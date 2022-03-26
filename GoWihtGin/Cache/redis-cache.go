package Cache

import (
	"GoWithGin/entity"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) MovieCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, movie *entity.Movie) {
	client := cache.getClient()

	// serialize Movie object to JSON
	json, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *entity.Movie {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	movie := entity.Movie{}
	err = json.Unmarshal([]byte(val), &movie)
	if err != nil {
		panic(err)
	}

	return &movie
}
