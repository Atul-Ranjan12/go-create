package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

// Pointer to the database connection
var dbCOnn = &DB{}

// setting up constants used by database
const MaxOpenDBConn = 10
const MaxIdleDBConn = 5
const MaxDBLifetime = 5 * time.Minute

// Function to connect to sql
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	// Set up database constants
	d.SetMaxOpenConns(MaxOpenDBConn)
	d.SetMaxIdleConns(MaxIdleDBConn)
	d.SetConnMaxLifetime(MaxDBLifetime)

	dbCOnn.SQL = d
	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbCOnn, err
}

// Tries to ping the database
func testDB(d *sql.DB) error {
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

// Creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	timeQuery := `
	SET TIMEZONE to 'UTC'
	`
	_, err = db.Exec(timeQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}
