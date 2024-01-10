package service

import (
	"context"
	"log/slog"
	"math/rand"
)

type iUrlRepository interface {
	SetUrl(ctx context.Context, url string, uri string) error
	GetUrl(ctx context.Context, uri string) (string, error)
}

type UrlService struct {
	repo iUrlRepository
}

// letterBytes is constant used for generate random uris
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewUrlService(repo iUrlRepository) *UrlService {
	return &UrlService{
		repo: repo,
	}
}

// SetUrl is function used to create URI, set them in redis and return URI
func (svc *UrlService) SetUrl(ctx context.Context, url string, lenght int) (string, error) {
	uri := randStringBytes(lenght)
	if err := svc.repo.SetUrl(ctx, url, uri); err != nil {
		slog.Error("filed to set URL : ", slog.String("err", err.Error()))
		return "", err
	}

	return uri, nil
}

func (svc *UrlService) GetUrl(ctx context.Context, uri string) (string, error) {
	url, err := svc.repo.GetUrl(ctx, uri)

	if err != nil {
		slog.Error("filed to get URL : ", slog.String("err", err.Error()))
		return "", err
	}

	return url, nil
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
