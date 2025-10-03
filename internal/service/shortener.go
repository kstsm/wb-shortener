package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gookit/slog"
	"github.com/kstsm/wb-shortener/internal/apperrors"
	"github.com/kstsm/wb-shortener/internal/dto"
	"github.com/kstsm/wb-shortener/internal/models"
)

func (s *Service) ShortenURL(ctx context.Context, req *dto.ShortenRequest) (*dto.ShortenResponse, error) {
	shortURL, err := s.resolveShortURL(ctx, req)
	if err != nil {
		return nil, err
	}

	link, err := s.repo.CreateLink(ctx, req.URL, shortURL, req.CustomAlias)
	if err != nil {
		return nil, fmt.Errorf("create link: %w", err)
	}

	if err := s.redis.SetLink(ctx, shortURL, link); err != nil {
		slog.Warn("Warning: failed to cache link: %v\n", err)
	}

	return &dto.ShortenResponse{
		ShortURL:    shortURL,
		OriginalURL: req.URL,
	}, nil
}

func (s *Service) Redirect(ctx context.Context, shortURL string, reqInfo models.RequestInfo) (string, error) {
	link, err := s.redis.GetLink(ctx, shortURL)
	if err != nil {
		link, err = s.repo.GetLinkByShortURL(ctx, shortURL)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return "", apperrors.ErrNotFound
			}

			return "", fmt.Errorf("get link from repo: %w", err)
		}

		if err := s.redis.SetLink(ctx, shortURL, link); err != nil {
			slog.Warn("failed to cache link: %v", err)
		}
	}

	if err := s.redis.IncrementClickCount(ctx, shortURL); err != nil {
		slog.Warn("failed to increment click count: %v", err)
	}

	if _, err := s.repo.CreateAnalytics(ctx, link.ID, reqInfo); err != nil {
		slog.Warn("failed to save analytics: %v", err)
	}

	return link.OriginalURL, nil
}

func (s *Service) GetAnalytics(ctx context.Context, shortURL string) (*dto.AnalyticsResponse, error) {
	link, err := s.repo.GetLinkByShortURL(ctx, shortURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("get link: %w", err)
	}

	totalClicks, err := s.repo.GetTotalClicks(ctx, link.ID)
	if err != nil {
		return nil, fmt.Errorf("get total clicks: %w", err)
	}

	dailyStats, err := s.repo.GetDailyStats(ctx, link.ID)
	if err != nil {
		return nil, fmt.Errorf("get daily stats: %w", err)
	}

	monthlyStats, err := s.repo.GetMonthlyStats(ctx, link.ID)
	if err != nil {
		return nil, fmt.Errorf("get monthly stats: %w", err)
	}

	userAgentStats, err := s.repo.GetUserAgentStats(ctx, link.ID)
	if err != nil {
		return nil, fmt.Errorf("get user agent stats: %w", err)
	}

	return &dto.AnalyticsResponse{
		ShortURL:       shortURL,
		OriginalURL:    link.OriginalURL,
		TotalClicks:    totalClicks,
		DailyStats:     mapDailyStats(dailyStats),
		MonthlyStats:   mapMonthlyStats(monthlyStats),
		UserAgentStats: mapUserAgentStats(userAgentStats),
	}, nil
}
