package database

import (
	"github.com/jinzhu/gorm"

	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is the DB that will performs all operation
type DB struct {
	DB *gorm.DB
}

// Open returns a DB connection
func Open(conn string) (*DB, error) {
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
