package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Db wraps a Gorm database connection object
type Db struct {
	// Db is the current database
	Db *gorm.DB
}

// newDb makes a new Db instance using the app config
// Returns:
// - *Db: New db instance
// - error: If one occurred while connecting to the db, nil on success
func newDb() (*Db, error) {
	// Load config
	config, err := config.LoadConfig()
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

// once is the thread safe object used ensure a singleton is only initialized 
// one time
var sync sync.Once

// instance holds the singleton Db instance
var db *Db

// GetDb returns the singleton instance of the Db struct.
// Returns:
//	- *Db: Singleton instance of Db
//	- error: If one occurred while initializing the first instance of the 
//		 singleton
func GetDb() (*Db, error) {
	once.Do(func() {
		db, err := 
		// TODO init singleton here
	})

	// TODO return singleton here
	// TODO convert to use this method
}
