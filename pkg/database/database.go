package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shashaneRanasinghe/simpleAPI/internal/config"
	"github.com/tryfix/log"
	"os"
)

type database struct {
	db *sql.DB
}

func NewDatabase() *database {
	return &database{}
}

func (d *database) InitDatabase() {

	cfg := config.DBConnection{
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		DBName:    os.Getenv("DB_NAME"),
		DBNetwork: os.Getenv("DB_NETWORK"),
	}
	// Get a database handle.
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", cfg.Username, cfg.Password, cfg.DBNetwork, cfg.DBHost,
		cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB ", err)
	}
	log.Info("Connected to Database")
	d.db = db
}

func (d *database) GetConnection() *sql.DB {
	return d.db
}
