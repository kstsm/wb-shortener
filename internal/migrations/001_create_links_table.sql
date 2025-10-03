-- +goose Up
CREATE TABLE IF NOT EXISTS links
(
    id           SERIAL PRIMARY KEY,
    original_url TEXT               NOT NULL,
    short_url    VARCHAR(20) UNIQUE NOT NULL,
    custom_alias VARCHAR(20) UNIQUE,
    created_at   TIMESTAMPTZ DEFAULT NOW(),
    updated_at   TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_links_short_url ON links (short_url);
CREATE INDEX IF NOT EXISTS idx_links_custom_alias ON links (custom_alias);
CREATE INDEX IF NOT EXISTS idx_links_created_at ON links (created_at);

-- +goose Down
DROP INDEX IF EXISTS idx_links_short_url;
DROP INDEX IF EXISTS idx_links_custom_alias;
DROP INDEX IF EXISTS idx_links_created_at;
DROP TABLE IF EXISTS links;