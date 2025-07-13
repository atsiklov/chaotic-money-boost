package db

import (
	"backend/internal/config"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context, cfg config.DbServer) *pgx.Conn {
	log.Println("🚀 Connecting to database...")
	dbConnection, err := pgx.Connect(ctx, cfg.DbConnectionParams())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	log.Println("Successfully connected to database ✅")
	return dbConnection
}
