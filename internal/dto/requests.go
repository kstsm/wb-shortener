package dto

type ShortenRequest struct {
	URL         string `json:"url"`
	CustomAlias string `json:"custom_alias,omitempty"`
}

type AnalyticsRequest struct {
	ShortURL string `json:"short_url"`
}
