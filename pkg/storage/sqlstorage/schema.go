package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"github.com/numary/go-libs/sharedlogging"
	"path"
)

type Schema interface {
	executor
	Initialize(ctx context.Context) error
	Table(name string) string
	Close(ctx context.Context) error
	BeginTx(ctx context.Context, s *sql.TxOptions) (*sql.Tx, error)
	Flavor() sqlbuilder.Flavor
	Name() string
}

type baseSchema struct {
	*sql.DB
	closeDb bool
	name    string
}

func (s *baseSchema) Name() string {
	return s.name
}

func (s *baseSchema) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	sharedlogging.GetLogger(ctx).Debugf("QueryContext: %s %s", query, args)
	return s.DB.QueryContext(ctx, query, args...)
}
func (s *baseSchema) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	sharedlogging.GetLogger(ctx).Debugf("QueryRowContext: %s %s", query, args)
	return s.DB.QueryRowContext(ctx, query, args...)
}
func (s *baseSchema) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	sharedlogging.GetLogger(ctx).Debugf("ExecContext: %s %s", query, args)
	return s.DB.ExecContext(ctx, query, args...)
}
func (s *baseSchema) Close(ctx context.Context) error {
	if s.closeDb {
		return s.DB.Close()
	}
	return nil
}

func (s *baseSchema) Table(name string) string {
	return name
}

func (s *baseSchema) Initialize(ctx context.Context) error {
	return nil
}

type PGSchema struct {
	baseSchema
	prefix string
}

func (s *PGSchema) Table(name string) string {
	return fmt.Sprintf(`"%s".%s`, s.prefix, name)
}

func (s *PGSchema) Initialize(ctx context.Context) error {
	_, err := s.ExecContext(ctx, fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS \"%s\"", s.name))
	return err
}

func (s *PGSchema) Flavor() sqlbuilder.Flavor {
	return sqlbuilder.PostgreSQL
}

type SQLiteSchema struct {
	baseSchema
}

func (s SQLiteSchema) Flavor() sqlbuilder.Flavor {
	return sqlbuilder.SQLite
}

type DB interface {
	Schema(ctx context.Context, name string) (Schema, error)
	Close(ctx context.Context) error
}

type postgresDB struct {
	db *sql.DB
}

func (p *postgresDB) Schema(ctx context.Context, name string) (Schema, error) {
	return &PGSchema{
		baseSchema: baseSchema{
			DB:   p.db,
			name: name,
		},
		prefix: name,
	}, nil
}

func (p *postgresDB) Close(ctx context.Context) error {
	return p.db.Close()
}

func NewPostgresDB(db *sql.DB) *postgresDB {
	return &postgresDB{
		db: db,
	}
}

type sqliteDB struct {
	directory string
	dbName    string
}

func (p *sqliteDB) Schema(ctx context.Context, name string) (Schema, error) {
	path := path.Join(
		p.directory,
		fmt.Sprintf("%s_%s.db", p.dbName, name),
	)
	db, err := OpenSQLDB(SQLite, path)
	if err != nil {
		return nil, err
	}

	return &SQLiteSchema{
		baseSchema: baseSchema{
			name:    name,
			DB:      db,
			closeDb: true,
		},
	}, nil
}

func (p *sqliteDB) Close(ctx context.Context) error {
	return nil
}

func NewSQLiteDB(directory, dbName string) *sqliteDB {
	return &sqliteDB{
		directory: directory,
		dbName:    dbName,
	}
}
