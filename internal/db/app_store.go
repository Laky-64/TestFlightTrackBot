package db

import (
	"github.com/Laky-64/TestFlightTrackBot/internal/db/models"
	"gorm.io/gorm"
	"maps"
	"slices"
)

type appStore struct {
	db *gorm.DB
}

func (ctx *appStore) Search(name string) ([]models.SearchResult, error) {
	var links []models.SearchResult
	if err := ctx.db.
		Raw(
			`
				WITH apps_fuzzy AS (
				    SELECT
				    	id,
				    	app_name,
				    	levenshtein(app_name, ?)::float / length(app_name) AS l_value
					FROM apps
					WHERE app_name IS NOT NULL
						AND levenshtein(app_name, ?)::float / length(app_name) <= 1.4
					LIMIT 3
				)
				SELECT
					a.id AS app_id,
					a.app_name,
					la.links_count,
					la.followers,
    				la.updated_at
				FROM apps_fuzzy a
				CROSS JOIN LATERAL (
					SELECT
						COUNT(l.*)    AS links_count,
						COUNT(cl.*) AS followers,
						MAX(l.updated_at) AS updated_at
					FROM links l
					LEFT JOIN chat_links cl ON cl.link_id = l.id
					WHERE l.app_id = a.id AND l.status IS NOT NULL
				) la
				ORDER BY a.l_value, la.followers DESC;
            `,
			name, name,
		).Scan(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

func (ctx *appStore) BulkUpsert(updates map[string]*uint) error {
	keys := slices.Collect(maps.Keys(updates))
	values := slices.Collect(maps.Values(updates))
	return bulkExec(
		ctx.db,
		`
			WITH
			input_data AS (
				SELECT *
				FROM UNNEST(
					ARRAY[?]::bigint[],
					ARRAY[?]
				) AS i(app_id, app_name)
			)
			MERGE INTO apps AS target
			USING input_data AS src ON (target.id = src.app_id OR target.app_name = src.app_name)
			WHEN MATCHED AND target.app_name IS DISTINCT FROM src.app_name THEN
				UPDATE SET app_name = src.app_name, updated_at = NOW()
			WHEN NOT MATCHED THEN
				INSERT (app_name, updated_at, created_at)
				VALUES (src.app_name, NOW(), NOW());
		`,
		values,
		keys,
	)
}
