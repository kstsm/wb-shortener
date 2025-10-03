package dto

type ShortenResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type AnalyticsResponse struct {
	ShortURL       string           `json:"short_url"`
	OriginalURL    string           `json:"original_url"`
	TotalClicks    int              `json:"total_clicks"`
	DailyStats     []DailyStats     `json:"daily_stats"`
	MonthlyStats   []MonthlyStats   `json:"monthly_stats"`
	UserAgentStats []UserAgentStats `json:"user_agent_stats"`
}

type DailyStats struct {
	Date   string `json:"date"`
	Clicks int    `json:"clicks"`
}

type MonthlyStats struct {
	Month  string `json:"month"`
	Clicks int    `json:"clicks"`
}

type UserAgentStats struct {
	UserAgent string `json:"user_agent"`
	Clicks    int    `json:"clicks"`
}
