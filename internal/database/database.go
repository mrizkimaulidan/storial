package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mrizkimaulidan/storial/internal/config"
)

type Database struct {
	c *config.Config
}

// Opening database connection.
// If something goes wrong, will throwing an fatal error.
func (d *Database) Open() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		d.c.DB_USERNAME,
		d.c.DB_PASSWORD,
		d.c.DB_HOST,
		d.c.DB_PORT,
		d.c.DB_DATABASE)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("error opening database", err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalln("error pinging to database", err)
	}

	return db
}

func NewDatabase() *Database {
	return &Database{
		c: config.New().GetConfig(),
	}
}
