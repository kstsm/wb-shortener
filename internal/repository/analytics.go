package repository

import (
	"context"
	"fmt"
	"github.com/kstsm/wb-shortener/internal/models"
	"time"
)

func (r *Repository) CreateAnalytics(ctx context.Context, linkID int, reqInfo models.RequestInfo) (*models.Analytics, error) {
	var analytics models.Analytics

	err := r.conn.QueryRow(ctx, CreateAnalyticsQuery,
		linkID,
		reqInfo.UserAgent,
		reqInfo.IP,
		reqInfo.Referer,
	).Scan(
		&analytics.ID,
		&analytics.LinkID,
		&analytics.UserAgent,
		&analytics.IPAddress,
		&analytics.Referer,
		&analytics.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("QueryRow-CreateAnalytics: %w", err)
	}

	return &analytics, nil
}

func (r *Repository) GetAnalyticsByLinkID(ctx context.Context, linkID int) ([]models.Analytics, error) {
	rows, err := r.conn.Query(ctx, GetAnalyticsByLinkIDQuery, linkID)
	if err != nil {
		return nil, fmt.Errorf("Query-GetAnalyticsByLinkID: %w", err)
	}
	defer rows.Close()

	var analytics []models.Analytics
	for rows.Next() {
		var a models.Analytics
		err := rows.Scan(
			&a.ID,
			&a.LinkID,
			&a.UserAgent,
			&a.IPAddress,
			&a.Referer,
			&a.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan analytics: %w", err)
		}
		analytics = append(analytics, a)
	}

	return analytics, nil
}

func (r *Repository) GetTotalClicks(ctx context.Context, linkID int) (int, error) {
	var totalClicks int
	err := r.conn.QueryRow(ctx, GetTotalClicksQuery, linkID).Scan(&totalClicks)
	if err != nil {
		return 0, fmt.Errorf("QueryRow-GetTotalClicks: %w", err)
	}

	return totalClicks, nil
}

func (r *Repository) GetDailyStats(ctx context.Context, linkID int) ([]models.DailyStats, error) {
	rows, err := r.conn.Query(ctx, GetDailyStatsQuery, linkID)
	if err != nil {
		return nil, fmt.Errorf("GetDailyStats-GetDailyStats: %w", err)
	}
	defer rows.Close()

	var stats []models.DailyStats
	for rows.Next() {
		var s models.DailyStats
		var date time.Time
		if err := rows.Scan(&date, &s.Clicks); err != nil {
			return nil, fmt.Errorf("scan daily stats: %w", err)
		}
		s.Date = date.Format("2006-01-02")
		stats = append(stats, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate daily stats rows: %w", err)
	}

	return stats, nil
}

func (r *Repository) GetUserAgentStats(ctx context.Context, linkID int) ([]models.UserAgentStats, error) {
	rows, err := r.conn.Query(ctx, GetUserAgentStatsQuery, linkID)
	if err != nil {
		return nil, fmt.Errorf("Query-GetUserAgentStats: %w", err)
	}
	defer rows.Close()

	var stats []models.UserAgentStats
	for rows.Next() {
		var s models.UserAgentStats
		if err := rows.Scan(&s.UserAgent, &s.Clicks); err != nil {
			return nil, fmt.Errorf("scan user-agent stats: %w", err)
		}
		stats = append(stats, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate user-agent stats rows: %w", err)
	}

	return stats, nil
}

func (r *Repository) GetMonthlyStats(ctx context.Context, linkID int) ([]models.MonthlyStats, error) {
	rows, err := r.conn.Query(ctx, GetMonthlyStatsQuery, linkID)
	if err != nil {
		return nil, fmt.Errorf("Query-GetMonthlyStats: %w", err)
	}
	defer rows.Close()

	var stats []models.MonthlyStats
	for rows.Next() {
		var s models.MonthlyStats
		var month time.Time
		if err := rows.Scan(&month, &s.Clicks); err != nil {
			return nil, fmt.Errorf("scan monthly stats: %w", err)
		}
		s.Month = month.Format("2006-01")
		stats = append(stats, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate monthly stats rows: %w", err)
	}

	return stats, nil
}
