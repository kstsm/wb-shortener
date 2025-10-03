package service

import (
	"context"

	"github.com/kstsm/wb-shortener/internal/cache"
	"github.com/kstsm/wb-shortener/internal/dto"
	"github.com/kstsm/wb-shortener/internal/models"
	"github.com/kstsm/wb-shortener/internal/repository"
)

type ServiceI interface {
	ShortenURL(ctx context.Context, req *dto.ShortenRequest) (*dto.ShortenResponse, error)
	Redirect(ctx context.Context, shortURL string, reqInfo models.RequestInfo) (string, error)
	GetAnalytics(ctx context.Context, shortURL string) (*dto.AnalyticsResponse, error)
}

type Service struct {
	repo  repository.RepositoryI
	redis cache.CacheI
}

func NewService(repo repository.RepositoryI, redisClient cache.CacheI) ServiceI {
	return &Service{
		repo:  repo,
		redis: redisClient,
	}
}
