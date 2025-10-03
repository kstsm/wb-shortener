package models

import (
	"net"
	"time"
)

type Analytics struct {
	ID        int       `json:"id"`
	LinkID    int       `json:"link_id"`
	UserAgent string    `json:"user_agent"`
	IPAddress net.IP    `json:"ip_address"`
	Referer   string    `json:"referer"`
	CreatedAt time.Time `json:"created_at"`
}

type AnalyticsResponse struct {
	ShortURL       string           `json:"short_url"`
	OriginalURL    string           `json:"original_url"`
	TotalClicks    int              `json:"total_clicks"`
	DailyStats     []DailyStats     `json:"daily_stats"`
	UserAgentStats []UserAgentStats `json:"user_agent_stats"`
}

type DailyStats struct {
	Date   string `json:"date"`
	Clicks int    `json:"clicks"`
}

type UserAgentStats struct {
	UserAgent string `json:"user_agent"`
	Clicks    int    `json:"clicks"`
}

type MonthlyStats struct {
	Month  string `json:"month"`
	Clicks int    `json:"clicks"`
}

type RequestInfo struct {
	UserAgent string
	IP        string
	Referer   string
}
