package main

import (
	"context"
	"database/sql"
)

const schema = `
CREATE TABLE IF NOT EXISTS registrations (
	id BIGSERIAL PRIMARY KEY,
	full_name TEXT NOT NULL,
	phone TEXT NOT NULL,
	discipline TEXT NOT NULL CHECK (discipline IN ('bmx', 'workout', 'trampoline')),
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
`

func migrate(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, schema)
	return err
}
