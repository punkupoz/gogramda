package resolver

import (
	"context"
	"fmt"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/punkupoz/gogramda/database"
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
	db *database.DB
}

// NewRoot : create a resolver root
func NewRoot() (*Resolver, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, dbuser, dbname, pass)
	db, err := database.Open(conn)
	if err != nil {
		return nil, err
	}

	db.DB.AutoMigrate(&User{})

	return &Resolver{db}, nil
}

// User : Resolver function for the "User" query
func (r *Resolver) User(ctx context.Context, args struct{ ID graphql.ID }) *UserResolver {
	p := &User{
		ID:        "1001",
		FirstName: "John",
		LastName:  "Doe",
	}
	return &UserResolver{p}
}
