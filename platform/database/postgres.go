package database

import (
	"fmt"
	"github.com/fabiotavarespr/go-env"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*sqlx.DB, error) {
	// Define database connection settings.
	maxConn := env.GetEnvWithDefaultAsInt("DB_MAX_CONNECTIONS", 100)
	maxIdleConn := env.GetEnvWithDefaultAsInt("DB_MAX_IDLE_CONNECTIONS", 10)
	maxLifetimeConn := env.GetEnvWithDefaultAsInt("DB_MAX_LIFETIME_CONNECTIONS", 2)

	// Define database connection for PostgreSQL.
	db, err := sqlx.Connect("pgx", env.GetEnvWithDefaultAsString("DB_SERVER_URL", "host=localhost port=5432 user=postgres password=123456 dbname=liveproject sslmode=disable"))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings.
	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
