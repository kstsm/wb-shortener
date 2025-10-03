package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/kstsm/wb-shortener/internal/config"
	"github.com/kstsm/wb-shortener/internal/models"
	"github.com/redis/go-redis/v9"
)

type CacheI interface {
	SetLink(ctx context.Context, shortURL string, link *models.Link) error
	GetLink(ctx context.Context, shortURL string) (*models.Link, error)
	IncrementClickCount(ctx context.Context, shortURL string) error
	GetClickCount(ctx context.Context, shortURL string) (int, error)
	Close() error
}

type Redis struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedis(ctx context.Context, cfg config.Config) *Redis {
	conf := &redis.Options{
		Addr: cfg.Redis.Address,
		DB:   cfg.Redis.DB,
	}

	rdb := redis.NewClient(conf)

	return &Redis{
		ctx:    ctx,
		client: rdb,
	}
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return r.client.Set(ctx, key, data, expiration).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", fmt.Errorf("key not found")
		}
		return "", fmt.Errorf("failed to get value: %w", err)
	}

	return result, nil
}

func (r *Redis) SetLink(ctx context.Context, shortURL string, link *models.Link) error {
	key := fmt.Sprintf("link:%s", shortURL)
	return r.Set(ctx, key, link, 0)
}

func (r *Redis) GetLink(ctx context.Context, shortURL string) (*models.Link, error) {
	key := fmt.Sprintf("link:%s", shortURL)
	data, err := r.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var link models.Link
	if err := json.Unmarshal([]byte(data), &link); err != nil {
		return nil, fmt.Errorf("failed to unmarshal link: %w", err)
	}

	return &link, nil
}

func (r *Redis) IncrementClickCount(ctx context.Context, shortURL string) error {
	key := fmt.Sprintf("clicks:%s", shortURL)
	return r.client.Incr(ctx, key).Err()
}

func (r *Redis) GetClickCount(ctx context.Context, shortURL string) (int, error) {
	key := fmt.Sprintf("clicks:%s", shortURL)
	result, err := r.client.Get(ctx, key).Int()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get click count: %w", err)
	}
	return result, nil
}

func (r *Redis) AddToPopularLinks(ctx context.Context, shortURL string, clickCount int) error {
	key := "popular_links"
	return r.client.ZAdd(ctx, key, redis.Z{
		Score:  float64(clickCount),
		Member: shortURL,
	}).Err()
}

func (r *Redis) GetPopularLinks(ctx context.Context, limit int) ([]string, error) {
	key := "popular_links"
	result, err := r.client.ZRevRange(ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get popular links: %w", err)
	}
	return result, nil
}

func (r *Redis) Close() error {
	return r.client.Close()
}
