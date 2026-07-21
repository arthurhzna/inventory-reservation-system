package database

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	host     string
	port     int
	username string
	password string
	dbName   string
	sslmode  string

	maxIdleConn    int
	maxOpenConn    int
	maxConnLifeMin int
}

func NewDatabase(
	host string,
	port int,
	username string,
	password string,
	dbName string,
	sslmode string,
	maxIdleConn int,
	maxOpenConn int,
	maxConnLifeMin int,
) *Database {

	return &Database{
		host:           host,
		port:           port,
		username:       username,
		password:       password,
		dbName:         dbName,
		sslmode:        sslmode,
		maxIdleConn:    maxIdleConn,
		maxOpenConn:    maxOpenConn,
		maxConnLifeMin: maxConnLifeMin,
	}
}

func (d *Database) Connect() (*sqlx.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		d.host,
		d.username,
		d.password,
		d.dbName,
		d.port,
		d.sslmode,
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(d.maxIdleConn)
	db.SetMaxOpenConns(d.maxOpenConn)
	db.SetConnMaxLifetime(
		time.Duration(d.maxConnLifeMin) * time.Minute,
	)

	return db, nil
}
