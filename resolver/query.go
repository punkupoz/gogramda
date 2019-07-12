package resolver

import (
	"context"
	graphql "github.com/graph-gophers/graphql-go"
)

type QueryResolver struct {}

func NewRoot() (*QueryResolver, error) {
	return &QueryResolver{}, nil
}

// Person : Resolver function for the "Person" query
func (r *QueryResolver) Person(ctx context.Context, args struct{ ID graphql.ID }) *PersonResolver {
	p := &Person{
		ID:        "1001",
		FirstName: "John",
		LastName:  "Doe",
	}
	return &PersonResolver{p}
}