package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/kstsm/wb-shortener/internal/apperrors"
	"github.com/kstsm/wb-shortener/internal/models"
)

func (r *Repository) CheckShortURLExists(ctx context.Context, shortURL string) (bool, error) {
	var exists bool
	err := r.conn.QueryRow(ctx, CheckShortURLExistsQuery, shortURL).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("QueryRow-CheckShortURLExists: %w", err)
	}

	return exists, nil
}

func (r *Repository) CreateLink(ctx context.Context, originalURL, shortURL, customAlias string) (*models.Link, error) {
	var link models.Link
	err := r.conn.QueryRow(ctx, CreateLinkQuery, originalURL, shortURL, customAlias).Scan(
		&link.ID,
		&link.OriginalURL,
		&link.ShortURL,
		&link.CustomAlias,
		&link.CreatedAt,
		&link.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("QueryRow-CreateLink: %w", err)
	}

	return &link, nil
}

func (r *Repository) GetLinkByShortURL(ctx context.Context, shortURL string) (*models.Link, error) {
	var link models.Link
	err := r.conn.QueryRow(ctx, GetLinkByShortURLQuery, shortURL).Scan(
		&link.ID,
		&link.OriginalURL,
		&link.ShortURL,
		&link.CustomAlias,
		&link.CreatedAt,
		&link.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("QueryRow-GetLinkByShortURL: %w", err)
	}

	return &link, nil
}

func (r *Repository) CheckCustomAliasExists(ctx context.Context, customAlias string) (bool, error) {
	var exists bool
	err := r.conn.QueryRow(ctx, CheckCustomAliasExistsQuery, customAlias).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("QueryRow-CheckCustomAliasExists: %w", err)
	}

	return exists, nil
}
