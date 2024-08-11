package database

import (
	"database/sql"

	"github.com/aofattaporn/go-cobra/configs"
	"github.com/aofattaporn/go-cobra/pkg/log"
	_ "github.com/go-sql-driver/mysql"
)

type IDatabase interface {
	GetDb() *sql.DB
	Close() error
}

type IQuery interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
}

func NewIQuery(db *sql.DB) IQuery { return db }

func NewIQueryTx(tx *sql.Tx) IQuery { return tx }

type mysqlDatabase struct {
	Db *sql.DB
}

func NewMysqlDatabase(cfg configs.IDbConfig, logger log.ILogger) (*sql.DB, error) {

	db, err := sql.Open("mysql", cfg.Url())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns())
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime())

	db.SetMaxOpenConns(cfg.MaxOpenConns())
	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime())

	return db, nil
}

func (m *mysqlDatabase) GetDb() *sql.DB {
	return m.Db
}

func (m *mysqlDatabase) Close() error {
	return m.Db.Close()
}
