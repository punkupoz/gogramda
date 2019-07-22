package resolver

import (
	"context"
	"fmt"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

func gqlIDP(id uint) graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return r
}

func gqlIDToUint(i graphql.ID) (uint, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(r), nil
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}

// ID resolves the ID field for Pet
func (p *UserResolver) ID(ctx context.Context) graphql.ID {
	return gqlIDP(p.p.ID)
}

func (r *UserResolver) FirstName() string {
	return r.p.FirstName
}

func (r *UserResolver) LastName() *string {
	return &r.p.LastName
}

type userCreateInput struct {
	ID        *graphql.ID
	FirstName *string
	LastName  *string
}

func (db *DB) createUser(ctx context.Context, input userCreateInput) (*User, error) {
	fmt.Print("createuserzzzzzfksdfjsdf")
	// Create user
	user := User{
		FirstName: *input.FirstName,
		LastName:  *input.LastName,
	}

	err := db.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type UserResolver struct {
	db *DB
	p  User
}

func (r *Resolver) CreateUser(ctx context.Context, args struct{ User userCreateInput }) (*UserResolver, error) {
	fmt.Print(os.Getenv("DB_HOST"))

	user, err := r.db.createUser(ctx, args.User)
	fmt.Print("asdas")
	if err != nil {
		return nil, err
	}

	s := UserResolver{
		db: r.db,
		p:  *user,
	}
	return &s, nil
}

func (db *DB) getUser(ctx context.Context, id uint) (*User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser resolves the getUser query
func (r *Resolver) GetUser(ctx context.Context, args struct{ ID graphql.ID }) (*UserResolver, error) {
	id, err := gqlIDToUint(args.ID)
	if err != nil {
		return nil, err
	}

	user, err := r.db.getUser(ctx, id)
	if err != nil {
		return nil, err
	}

	s := UserResolver{
		db: r.db,
		p:  *user,
	}

	return &s, nil
}
