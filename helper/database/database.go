package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Err error

// Creates a connection to the db
func ConnectDb() {
	// Create Data Source Name for mysql Driver by taking user,password,host,ip,database name from env file
	dsn := os.Getenv("DSN")
	// Establishes a connection to the database named in the dsn using gorm
	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Handling error
	if Err != nil {
		os.Exit(1)
	}
}

// Closes the connection of the db
func CloseDb() {
	mysqlDb, _ := DB.DB()
	mysqlDb.Close()
}
