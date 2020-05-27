// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const listBar = `-- name: ListBar :many
SELECT baz FROM foo.bar
`

func (q *Queries) ListBar(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listBar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var baz string
		if err := rows.Scan(&baz); err != nil {
			return nil, err
		}
		items = append(items, baz)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
