package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"go-wallet/internal/shared/src/infrastructure/database"

	"gorm.io/gorm"
)

func SetupPostgres(t *testing.T) *gorm.DB {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	container, err := postgres.Run(
		ctx,
		"postgres:17-alpine",
		postgres.WithDatabase("go_wallet_test"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Fatalf("could not start postgres container: %v", err)
	}

	t.Cleanup(func() {
		terminateCtx, terminateCancel := context.WithTimeout(
			context.Background(),
			30*time.Second,
		)
		defer terminateCancel()

		if err := container.Terminate(terminateCtx); err != nil {
			t.Errorf("could not terminate postgres container: %v", err)
		}
	})

	dsn, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("could not get postgres connection string: %v", err)
	}

	db, err := database.NewConnection(dsn, database.PoolConfig{
		MaxOpenConnections: 5,
		MaxIdleConnections: 2,
		MaxLifetime:        5 * time.Minute,
		MaxIdleTime:        1 * time.Minute,
	})
	if err != nil {
		t.Fatalf("could not connect to postgres: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("could not get sql.DB: %v", err)
	}

	t.Cleanup(func() {
		if err := sqlDB.Close(); err != nil {
			t.Errorf("could not close database connection: %v", err)
		}
	})

	return db
}
