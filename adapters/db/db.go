package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Adapter{db}, err
}

func (adap *Adapter) Close() {
	if err := adap.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func (adap *Adapter) ToHistory(answer int32, operation string) error {
	query := `INSERT INTO arith_history VALUES (?, ?, ?);`
	_, err := adap.db.Exec(query, time.Now(), answer, operation)
	return err
}
