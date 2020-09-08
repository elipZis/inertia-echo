package repository

import (
	"elipzis.com/inertia-echo/repository/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strconv"
)

//
type Database struct {
	Conn *gorm.DB
}

// A global reference to this database instance for later defer.close
var DB *Database

// Create an initial connection to the database and cache here
func NewDatabase() (this *Database) {
	this = new(Database)
	conn, err := gorm.Open(os.Getenv("DB_CONNECTION"), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSL")))
	if err != nil {
		log.Fatal("[Database] Error while connecting to the database: ", err)
	}
	this.Conn = conn
	DB = this

	// Debug output?
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	this.Conn.LogMode(debug)

	// Migrate (if anything)
	this.AutoMigrate()

	return this
}

//
func (this *Database) AutoMigrate() {
	this.Conn.AutoMigrate(
		&model.Contact{},
		&model.Organization{},
		&model.User{},
	)
}
