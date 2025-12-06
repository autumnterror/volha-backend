package redis

import (
	"context"
	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/internal/gateway/config"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type Client struct {
	Rdb *redis.Client
	m   *sync.Mutex
}

func MustNew(cfg *config.Config) *Client {
	if c, err := newClient(cfg); err != nil {
		log.Panic(err)
		return nil
	} else {
		return c
	}
}

func newClient(cfg *config.Config) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPw,
		DB:       0,
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second)
	defer done()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	var mutex sync.Mutex
	return &Client{Rdb: rdb, m: &mutex}, nil
}
