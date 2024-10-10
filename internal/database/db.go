package database

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shlmvgleb/em-task/internal/config"
)

func New(config *config.PostgresConfig, ctx context.Context) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("error while creating a connection to database: %s", err)
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		log.Fatalf("error while ping a database: %s", err)
		return nil, err
	}

	log.Infoln("Successfully connected to postgres")
	return db, nil
}
