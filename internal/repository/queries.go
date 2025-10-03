package repository

const (
	CreateLinkQuery = `
INSERT INTO links (original_url, 
                   short_url, 
                   custom_alias)
VALUES ($1, $2, NULLIF($3, ''))
RETURNING id, original_url, short_url, custom_alias, created_at, updated_at;
`

	GetLinkByShortURLQuery = `
SELECT id, 
       original_url, 
	   short_url, 
	   custom_alias, 
	   created_at, 
	   updated_at 
FROM links 
WHERE short_url = $1
`
	CreateAnalyticsQuery = `
INSERT INTO analytics (
    link_id,
    user_agent,
    ip_address,
    referer
) 
VALUES ($1, $2, $3, $4) 
RETURNING id, link_id, user_agent, ip_address, referer, created_at;
`
	CheckShortURLExistsQuery = `
SELECT EXISTS(SELECT 1 FROM links WHERE short_url = $1)
`

	CheckCustomAliasExistsQuery = `
SELECT EXISTS(SELECT 1 FROM links WHERE custom_alias = $1)
`

	GetAnalyticsByLinkIDQuery = `
SELECT a.id, 
       a.link_id, 
       a.user_agent, 
       a.ip_address, 
       a.referer, 
       a.created_at
FROM analytics a
WHERE a.link_id = $1
ORDER BY a.created_at DESC
`

	GetTotalClicksQuery = `
SELECT COUNT(*) FROM analytics WHERE link_id = $1
`

	GetDailyStatsQuery = `
SELECT DATE(created_at) AS date, COUNT(*) AS clicks
FROM analytics 
WHERE link_id = $1 
GROUP BY DATE(created_at)
ORDER BY date DESC
`

	GetUserAgentStatsQuery = `
SELECT user_agent, COUNT(*) AS clicks
FROM analytics 
WHERE link_id = $1 AND user_agent IS NOT NULL
GROUP BY user_agent
ORDER BY clicks DESC
LIMIT 10
`

	GetMonthlyStatsQuery = `
SELECT DATE_TRUNC('month', created_at) AS month, COUNT(*) AS clicks
FROM analytics 
WHERE link_id = $1 
GROUP BY DATE_TRUNC('month', created_at)
ORDER BY month DESC
`
)
