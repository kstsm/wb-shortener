package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kstsm/wb-shortener/internal/models"
)

type RepositoryI interface {
	CreateLink(ctx context.Context, originalURL, shortURL, customAlias string) (*models.Link, error)
	GetLinkByShortURL(ctx context.Context, shortURL string) (*models.Link, error)
	CheckShortURLExists(ctx context.Context, shortURL string) (bool, error)
	CheckCustomAliasExists(ctx context.Context, customAlias string) (bool, error)
	CreateAnalytics(ctx context.Context, linkID int, reqInfo models.RequestInfo) (*models.Analytics, error)
	GetTotalClicks(ctx context.Context, linkID int) (int, error)
	GetDailyStats(ctx context.Context, linkID int) ([]models.DailyStats, error)
	GetMonthlyStats(ctx context.Context, linkID int) ([]models.MonthlyStats, error)
	GetUserAgentStats(ctx context.Context, linkID int) ([]models.UserAgentStats, error)
}

type Repository struct {
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) RepositoryI {
	return &Repository{
		conn: conn,
	}
}
