-- +goose Up
CREATE TABLE IF NOT EXISTS analytics
(
    id         SERIAL PRIMARY KEY,
    link_id    INT NOT NULL REFERENCES links (id) ON DELETE CASCADE,
    user_agent TEXT,
    ip_address INET,
    referer    TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_analytics_link_id ON analytics (link_id);
CREATE INDEX IF NOT EXISTS idx_analytics_created_at ON analytics (created_at);
CREATE INDEX IF NOT EXISTS idx_analytics_user_agent ON analytics (user_agent);

-- +goose Down
DROP INDEX IF EXISTS idx_analytics_link_id;
DROP INDEX IF EXISTS idx_analytics_created_at;
DROP INDEX IF EXISTS idx_analytics_user_agent;
DROP TABLE IF EXISTS analytics;