package simbad

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"server/internal/client/kv"
)

const cacheKeyPrefix = "simbad:"

type querier interface {
	QueryObject(ctx context.Context, identifier string) (*ObjectInfo, error)
}

type CachedClient struct {
	inner querier
	kv    kv.Client
	ttl   int
}

func NewCachedClient(inner querier, kvClient kv.Client, ttlSeconds int) *CachedClient {
	return &CachedClient{
		inner: inner,
		kv:    kvClient,
		ttl:   ttlSeconds,
	}
}

func (c *CachedClient) QueryObject(ctx context.Context, identifier string) (*ObjectInfo, error) {
	key := normalizeCacheKey(identifier)

	data, found, err := c.kv.Get(ctx, key)
	if err != nil {
		log.Printf("cache get error for %q: %v", key, err)
	} else if found {
		var info ObjectInfo
		if err := json.Unmarshal(data, &info); err != nil {
			log.Printf("cache unmarshal error for %q: %v", key, err)
		} else {
			return &info, nil
		}
	}

	info, err := c.inner.QueryObject(ctx, identifier)
	if err != nil {
		return nil, err
	}

	go c.cacheResult(key, info)

	return info, nil
}

func (c *CachedClient) cacheResult(key string, info *ObjectInfo) {
	data, err := json.Marshal(info)
	if err != nil {
		log.Printf("cache marshal error for %q: %v", key, err)
		return
	}

	if err := c.kv.Put(context.Background(), key, data, c.ttl); err != nil {
		log.Printf("cache put error for %q: %v", key, err)
	}
}

func normalizeCacheKey(identifier string) string {
	normalized := strings.ToLower(identifier)
	normalized = strings.ReplaceAll(normalized, " ", "")
	return cacheKeyPrefix + normalized
}
