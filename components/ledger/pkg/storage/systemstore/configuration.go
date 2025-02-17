package systemstore

import (
	"context"

	"github.com/formancehq/ledger/pkg/core"
	storageerrors "github.com/formancehq/ledger/pkg/storage/errors"
	"github.com/uptrace/bun"
)

const configTableName = "configuration"

type configuration struct {
	bun.BaseModel `bun:"configuration,alias:configuration"`

	Key     string    `bun:"key,type:varchar(255),pk"` // Primary key
	Value   string    `bun:"value,type:text"`
	AddedAt core.Time `bun:"addedAt,type:timestamp"`
}

func (s *Store) CreateConfigurationTable(ctx context.Context) error {
	_, err := s.schema.NewCreateTable(configTableName).
		Model((*configuration)(nil)).
		IfNotExists().
		Exec(ctx)

	return storageerrors.PostgresError(err)
}

func (s *Store) GetConfiguration(ctx context.Context, key string) (string, error) {
	query := s.schema.NewSelect(configTableName).
		Model((*configuration)(nil)).
		Column("value").
		Where("key = ?", key).
		Limit(1).
		String()

	row := s.schema.QueryRowContext(ctx, query)
	if row.Err() != nil {
		return "", storageerrors.PostgresError(row.Err())
	}
	var value string
	if err := row.Scan(&value); err != nil {
		return "", storageerrors.PostgresError(err)
	}

	return value, nil
}

func (s *Store) InsertConfiguration(ctx context.Context, key, value string) error {
	config := &configuration{
		Key:     key,
		Value:   value,
		AddedAt: core.Now(),
	}

	_, err := s.schema.NewInsert(configTableName).
		Model(config).
		Exec(ctx)

	return storageerrors.PostgresError(err)
}
