package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type User struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}

type UserResolver struct {
	p *User
}

func (r *UserResolver) ID() graphql.ID {
	return r.p.ID
}

func (r *UserResolver) FirstName() string {
	return r.p.FirstName
}

func (r *UserResolver) LastName() *string {
	return &r.p.LastName
}
