package resolver

import (
	"context"
	"fmt"
	graphql "github.com/graph-gophers/graphql-go"
)

type User struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}

type UserResolver struct {
	db *DB
	p  User
}

func (r *UserResolver) ID() *graphql.ID {
	return &r.p.ID
}

func (r *UserResolver) FirstName() *string {
	return &r.p.FirstName
}

func (r *UserResolver) LastName() *string {
	return &r.p.LastName
}

type userCreateInput struct {
	FirstName string
	LastName  string
}

func (db *DB) createUser(ctx context.Context, input userCreateInput) (*User, error) {
	// Create user
	user := User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err := db.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Resolver) CreateUser(ctx context.Context, args struct{ User userCreateInput }) (*UserResolver, error) {
	user, err := r.db.createUser(ctx, args.User)
	if err != nil {
		return nil, err
	}

	s := UserResolver{
		db: r.db,
		p:  *user,
	}
	return &s, nil
}
