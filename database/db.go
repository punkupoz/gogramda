package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	host   = os.Getenv("DB_HOST")
	port   = os.Getenv("DB_PORT")
	user   = os.Getenv("DB_USER")
	dbname = os.Getenv("DB_NAME")
	pass   = os.Getenv("DB_PASS")
)

// DB is the DB that will performs all operation
type DB struct {
	DB *gorm.DB
}

// NewDB returns a new DB connection
func newDB(path string) (*DB, error) {
	// open a new connection
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, pass)
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
