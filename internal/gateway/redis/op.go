package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/autumnterror/breezynotes/pkg/utils/format"
	"github.com/autumnterror/volha-backend/api/proto/gen"

	"github.com/redis/go-redis/v9"
	"time"
)

const cacheKey = "dictionaries:all"

var CashedCategoriesId []string

func (c *Client) GetDictionaries() (*productsRPC.Dictionaries, error) {
	const (
		op = "redis.GetAllDictionaries"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ds, err := c.Rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var cached productsRPC.Dictionaries
		if err := json.Unmarshal([]byte(ds), &cached); err != nil {
			return nil, format.Error(op, err)
		}
		return &cached, nil
	}
	if err == redis.Nil {
		return nil, format.Error(op, fmt.Errorf("no cache"))
	}

	return nil, format.Error(op, err)
}
func (c *Client) SetDictionaries(d *productsRPC.Dictionaries) error {
	const (
		op = "redis.SetDictionaries"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	bytes, err := json.Marshal(d)
	if err != nil {
		return format.Error(op, fmt.Errorf("SetDictionariesToCache: failed to marshal: %w", err))
	}

	if err := c.Rdb.Set(ctx, cacheKey, bytes, 0).Err(); err != nil {
		return format.Error(op, err)
	}

	return nil
}

func (c *Client) CleanDictionaries() error {
	const (
		op = "redis.CleanDictionaries"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := c.Rdb.Del(ctx, cacheKey).Err(); err != nil {
		return format.Error(op, err)
	}
	for _, id := range CashedCategoriesId {
		if err := c.Rdb.Del(ctx, cacheKey+id).Err(); err != nil {
			return format.Error(op, err)
		}
	}
	CashedCategoriesId = []string{}
	return nil
}

func (c *Client) GetDictionariesByCategory(id string) (*productsRPC.Dictionaries, error) {
	const (
		op = "redis.GetDictionariesByCategory"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ds, err := c.Rdb.Get(ctx, cacheKey+id).Result()
	if err == nil {
		var cached productsRPC.Dictionaries
		if err := json.Unmarshal([]byte(ds), &cached); err != nil {
			return nil, format.Error(op, err)
		}
		return &cached, nil
	}
	if err == redis.Nil {
		return nil, format.Error(op, fmt.Errorf("no cache"))
	}

	return nil, format.Error(op, err)
}
func (c *Client) SetDictionariesByCategory(id string, d *productsRPC.Dictionaries) error {
	const (
		op = "redis.SetDictionariesByCategory"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	bytes, err := json.Marshal(d)
	if err != nil {
		return format.Error(op, fmt.Errorf("SetDictionariesToCache: failed to marshal: %w", err))
	}

	if err := c.Rdb.Set(ctx, cacheKey+id, bytes, 0).Err(); err != nil {
		return format.Error(op, err)
	}
	c.m.Lock()
	CashedCategoriesId = append(CashedCategoriesId, id)
	c.m.Unlock()

	return nil
}
