package sqlite

import (
	"github.com/AmitKarnam/KeyCloak/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Asserts that the implementing type DB implements the GetConnection method
var _ database.DB = &sqliteDB{}

// Singleton Pattern is implemented using this reference variable, which is used by other modules to access this singleton reference.
var DB database.DB

// Defines the implementing type DB, which in this case is sqlite
type sqliteDB struct {
	path string
}

// Assigns the singleton reference variable the sqlite DB instance
func DBInit(path string) {
	DB = New(path)
}

// Creates a new sqliteDB instance
func New(path string) *sqliteDB {
	return &sqliteDB{
		path: path,
	}
}

// Implementation of the GetConnection method
func (s *sqliteDB) GetConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(s.path), &gorm.Config{
		SkipDefaultTransaction: true,
	})
}
