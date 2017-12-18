package db

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Db wraps a Gorm database connection object
type Db struct {
	// Db is the current database
	Db *gorm.DB
}

// NewDb makes a new Db instance using the app config
// Returns:
// - *Db: New db instance
// - error: If one occurred while connecting to the db, nil on success
func NewDb() (*Db, error) {
	// Get config
	config, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading app config: %s", err.Error())
	}

	// Connect
	connStr := fmt.Sprintf("host=%s, user=%s, dbname=%s, sslmode=disable, password=%s",
		config.Db.Host,
		config.Db.User,
		config.Db.Db,
		config.Db.Password)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to db: %s", err.Error())
	}

	// Make Db struct
	dbStruct = Db{db}

	return dbStruct
}
