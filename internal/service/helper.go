package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/kstsm/wb-shortener/internal/apperrors"
	"github.com/kstsm/wb-shortener/internal/dto"
	"github.com/kstsm/wb-shortener/internal/models"
	"math/big"
)

const (
	maxAttempts    = 20
	shortURLLength = 8
	base62Chars    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func (s *Service) resolveShortURL(ctx context.Context, req *dto.ShortenRequest) (string, error) {
	if req.CustomAlias != "" {
		exists, err := s.repo.CheckCustomAliasExists(ctx, req.CustomAlias)
		if err != nil {
			return "", fmt.Errorf("check custom alias: %w", err)
		}
		if exists {
			return "", apperrors.ErrAliasAlreadyExists
		}

		return req.CustomAlias, nil
	}

	alias, err := s.generateShortURL(ctx)
	if err != nil {
		return "", fmt.Errorf("generate short URL: %w", err)
	}

	return alias, nil
}

func (s *Service) generateShortURL(ctx context.Context) (string, error) {
	for i := 0; i < maxAttempts; i++ {
		shortURL, err := generateRandomBase62(shortURLLength)
		if err != nil {
			return "", fmt.Errorf("generate random short URL: %w", err)
		}

		exists, err := s.repo.CheckShortURLExists(ctx, shortURL)
		if err != nil {
			return "", fmt.Errorf("check short URL existence: %w", err)
		}

		if !exists {
			return shortURL, nil
		}
	}

	return "", fmt.Errorf("failed to generate unique short URL after %d attempts", maxAttempts)
}

func generateRandomBase62(length int) (string, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(base62Chars))))
		if err != nil {
			return "", err
		}
		result[i] = base62Chars[num.Int64()]
	}

	return string(result), nil
}

func mapDailyStats(stats []models.DailyStats) []dto.DailyStats {
	result := make([]dto.DailyStats, len(stats))
	for i, s := range stats {
		result[i] = dto.DailyStats{
			Date:   s.Date,
			Clicks: s.Clicks,
		}
	}

	return result
}

func mapMonthlyStats(stats []models.MonthlyStats) []dto.MonthlyStats {
	result := make([]dto.MonthlyStats, len(stats))
	for i, s := range stats {
		result[i] = dto.MonthlyStats{
			Month:  s.Month,
			Clicks: s.Clicks,
		}
	}

	return result
}

func mapUserAgentStats(stats []models.UserAgentStats) []dto.UserAgentStats {
	result := make([]dto.UserAgentStats, len(stats))
	for i, s := range stats {
		result[i] = dto.UserAgentStats{
			UserAgent: s.UserAgent,
			Clicks:    s.Clicks,
		}
	}

	return result
}
