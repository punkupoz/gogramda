package resolver

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type QueryResolver struct{}

func NewRoot() (*QueryResolver, error) {
	return &QueryResolver{}, nil
}

// User : Resolver function for the "User" query
func (r *QueryResolver) User(ctx context.Context, args struct{ ID graphql.ID }) *UserResolver {
	p := &User{
		ID:        "1001",
		FirstName: "John",
		LastName:  "Doe",
	}
	return &UserResolver{p}
}
