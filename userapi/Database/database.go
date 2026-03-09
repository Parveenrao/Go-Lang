package database

import (
	"fmt"
	"time"
	"userapi/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB takes a DBConfig and returns a ready-to-use *gorm.DB.
//
// Why return *gorm.DB instead of storing a global?
//   - The caller (main.go) owns the DB and passes it down
//   - Easy to create a second DB connection (e.g. read replica)
//   - Easy to pass a test DB in unit tests
func NewDB(cfg config.DBConfig) (*gorm.DB, error) {
	// Build the DSN (Data Source Name) from individual config fields.
	// This is safer than a raw hardcoded string — fields come from env vars.
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	// gorm.Config lets us customize GORM's behaviour.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// logger.Info prints every SQL query to stdout.
		// Use logger.Silent in production to avoid log noise.
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		// Wrap the error with context so callers know where it came from.
		return nil, fmt.Errorf("database.NewDB: failed to connect: %w", err)
	}

	// -------------------------------------------------------
	// Connection Pool Configuration — critical for production.
	// Without this, every request opens a new DB connection,
	// which will crash your database under any real load.
	// -------------------------------------------------------
	sqlDB, err := db.DB() // get the underlying *sql.DB
	if err != nil {
		return nil, fmt.Errorf("database.NewDB: failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)               // max simultaneous connections to Postgres
	sqlDB.SetMaxIdleConns(10)               // connections kept alive when idle
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // recycle connections after 5 min

	return db, nil
}