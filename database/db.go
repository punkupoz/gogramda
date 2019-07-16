package database

import (
	"github.com/jinzhu/gorm"
	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/punkupoz/gogramda/resolver"
)

// DB is the DB that will performs all operation
type DB struct {
	DB *gorm.DB
}

// NewDB returns a new DB connection
func newDB(path string) (*DB, error) {
	// connect to the example db, create if it doesn't exist
	db, err := gorm.Open("postgres", path)
	if err != nil {
		return nil, err
	}

	// drop tables and all data, and recreate them fresh for this run
	db.DropTableIfExists(&resolver.User{})
	db.AutoMigrate(&resolver.User{})

	// put all the users into the db
	for _, u := range users {
		if err := db.Create(&u).Error; err != nil {
			return nil, err
		}
	}

	return &DB{db}, nil
}

// TEST DATA TO BE PUT INTO THE DB
var users = []resolver.User{
	resolver.User{FirstName: "Alice"},
	resolver.User{FirstName: "Bob"},
	resolver.User{FirstName: "Charlie"},
}
