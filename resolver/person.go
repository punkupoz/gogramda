package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type Person struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}


type PersonResolver struct {
	p *Person
}

func (r *PersonResolver) ID() graphql.ID {
	return r.p.ID
}

func (r *PersonResolver) FirstName() string {
	return r.p.FirstName
}

func (r *PersonResolver) LastName() *string {
	return &r.p.LastName
}

