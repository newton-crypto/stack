package _17_optimized_segments

import (
	"context"
	"testing"

	"github.com/formancehq/ledger/pkg/core"
	ledgerstore "github.com/formancehq/ledger/pkg/storage/sqlstorage/ledger"
	"github.com/formancehq/ledger/pkg/storage/sqlstorage/migrations"
	"github.com/formancehq/ledger/pkg/storage/sqlstorage/sqlstoragetesting"
	"github.com/formancehq/stack/libs/go-libs/pgtesting"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/require"
)

type Transaction struct {
	ledgerstore.Transaction
	ID uint64 `bun:"id,type:bigint,unique"`
}

func TestMigrate17(t *testing.T) {
	require.NoError(t, pgtesting.CreatePostgresServer())
	defer func() {
		require.NoError(t, pgtesting.DestroyPostgresServer())
	}()

	driver := sqlstoragetesting.StorageDriver(t)

	require.NoError(t, driver.Initialize(context.Background()))
	store, _, err := driver.GetLedgerStore(context.Background(), uuid.New(), true)
	require.NoError(t, err)

	schema := store.(*ledgerstore.Store).Schema()

	ms, err := migrations.CollectMigrationFiles(ledgerstore.MigrationsFS)
	require.NoError(t, err)

	modified, err := migrations.Migrate(context.Background(), schema, ms[0:17]...)
	require.NoError(t, err)
	require.True(t, modified)

	now := core.Now()

	tr := &Transaction{
		ID: 0,
		Transaction: ledgerstore.Transaction{
			Timestamp: now,
			Postings: []byte(`[
				{"source": "world", "destination": "users:001", "asset": "USD", "amount": 100}
			]`),
			Metadata: []byte(`{}`),
		},
	}
	_, err = schema.NewInsert(ledgerstore.TransactionsTableName).
		Column("id", "timestamp", "postings", "metadata").
		Model(tr).
		Exec(context.Background())
	require.NoError(t, err)

	modified, err = migrations.Migrate(context.Background(), schema, ms[17])
	require.NoError(t, err)
	require.True(t, modified)

	sb := schema.NewSelect("postings").
		Model((*ledgerstore.Postings)(nil)).
		Column("txid", "posting_index", "source", "destination").
		Where("txid = 0")

	row := store.(*ledgerstore.Store).Schema().QueryRowContext(context.Background(), sb.String())
	require.NoError(t, row.Err())

	var txid uint64
	var postingIndex int
	var source, destination string
	require.NoError(t, err, row.Scan(&txid, &postingIndex, &source, &destination))
	require.Equal(t, uint64(0), txid)
	require.Equal(t, 0, postingIndex)
	require.Equal(t, `["world"]`, source)
	require.Equal(t, `["users", "001"]`, destination)
}
