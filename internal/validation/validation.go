package validation

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

const (
	MinShortURLLength = 3
	MaxShortURLLength = 15
)

var (
	ErrEmpty         = errors.New("value cannot be empty")
	ErrTooShort      = errors.New("value is too short")
	ErrTooLong       = errors.New("value is too long")
	ErrInvalid       = errors.New("value contains invalid characters or format")
	ErrInvalidScheme = errors.New("only http/https are allowed")
	ErrMissingHost   = errors.New("URL must include host")
)

var shortURLRegex = regexp.MustCompile(`^[0-9a-zA-Z_-]+$`)

func IsValidURL(rawURL string) error {
	if strings.TrimSpace(rawURL) == "" {
		return ErrEmpty
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ErrInvalid
	}

	if parsedURL.Scheme == "" {
		return ErrInvalid
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return ErrInvalidScheme
	}

	if strings.TrimSpace(parsedURL.Host) == "" {
		return ErrMissingHost
	}

	return nil
}

func IsValidShortURL(shortURL string) error {
	if strings.TrimSpace(shortURL) == "" {
		return ErrEmpty
	}

	if !shortURLRegex.MatchString(shortURL) {
		return ErrInvalid
	}

	if len(shortURL) < MinShortURLLength {
		return ErrTooShort
	}

	if len(shortURL) > MaxShortURLLength {
		return ErrTooLong
	}

	return nil
}
