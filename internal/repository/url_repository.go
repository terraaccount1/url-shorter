package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type UrlRepository struct {
	redisDB *redis.Client
}

func NewUrlRepository(redisDB *redis.Client) *UrlRepository {
	return &UrlRepository{
		redisDB: redisDB,
	}
}

func (repo *UrlRepository) SetUrl(ctx context.Context, url string, uri string) error {
	err := repo.redisDB.Set(ctx, uri, url, time.Minute*10).Err()
	if err != redis.Nil {
		return err
	}

	return nil
}

func (repo *UrlRepository) GetUrl(ctx context.Context, uri string) (string, error) {
	url, err := repo.redisDB.Get(ctx, uri).Result()
	if err != nil {
		return "", err
	}
	slog.Info(url)
	return url, nil
}
