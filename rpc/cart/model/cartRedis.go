package model

import "github.com/tal-tech/go-zero/core/stores/redis"

const (
	connStr = "localhost:6379"
)

type (
	CartRedis interface {
		DelOneByKey(key string) error
	}

	defaultCarRedis struct {
		redis.Redis
	}
)

func NewCartRedis() CartRedis {
	return &defaultCarRedis{
		Redis: *redis.NewRedis(connStr, redis.NodeType, ""),
	}
}

func (r *defaultCarRedis) DelOneByKey(key string) error {
	_, err := r.Del(key)
	if err != nil {
		return err
	}
	return nil
}
