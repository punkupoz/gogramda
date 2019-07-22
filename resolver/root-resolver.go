package resolver

import (
	"fmt"
	"os"
)

var (
	host   = os.Getenv("DB_HOST")
	port   = os.Getenv("DB_PORT")
	dbuser = os.Getenv("DB_USER")
	dbname = os.Getenv("DB_NAME")
	pass   = os.Getenv("DB_PASS")
)

// Resolver : connection
type Resolver struct {
	db *DB
}

// NewRoot : create a resolver root
func NewRoot() (*Resolver, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, dbuser, dbname, pass)
	db, err := Open(conn)
	if err != nil {
		return nil, err
	}

	db.DB.AutoMigrate(&User{})

	return &Resolver{db}, nil
}
