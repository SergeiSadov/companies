package db

import (
	"database/sql"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	GetConn() (conn *gorm.DB, err error)
	GetUnderlyingConn() (conn *sql.DB, err error)
	Close() (err error)
}

type Database struct {
	driver string
	dsn    string
	conn   *gorm.DB
}

func New(
	driver string,
	dsn string,
) (*Database, error) {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{
		driver: driver,
		dsn:    dsn,
		conn:   conn,
	}, nil
}

func (db *Database) GetConn() (conn *gorm.DB, err error) {
	if db.conn == nil {
		return conn, errors.New("empty conn")
	}

	return db.conn, nil
}

func (db *Database) GetUnderlyingConn() (conn *sql.DB, err error) {
	conn, err = db.conn.DB()
	if err != nil {
		return
	}

	return
}

func (db *Database) Close() (err error) {
	database, err := db.conn.DB()
	if err != nil {
		return err
	}

	return database.Close()
}
