package models

import (
	"time"
)

type Link struct {
	ID          int       `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CustomAlias *string   `json:"custom_alias,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateLinkRequest struct {
	URL         string `json:"url"`
	CustomAlias string `json:"custom_alias,omitempty"`
}

type CreateLinkResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
