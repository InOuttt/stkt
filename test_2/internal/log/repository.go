package log

import (
	"database/sql"

	"github.com/inouttt/stkt/test_2/internal/config"
	"github.com/inouttt/stkt/test_2/internal/entity"
)

const (
	QueryInsert = `INSERT INTO ` + entity.LogTableName + `
		(request_body, response_body, created_at) 
		VALUES (?, ?, ?)`
)

type Repository interface {
	Create(req entity.Log) error
}

type log struct {
	Db *config.MysqlSession
}

func NewRepository(db *config.MysqlSession) Repository {
	return &log{
		Db: db,
	}
}

func (l log) insertRecord(tx *sql.Tx, req entity.Log) error {
	stmt, err := tx.Prepare(QueryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(req.ToInsertQry()...); err != nil {
		return err
	}

	return nil
}
func (l *log) Create(req entity.Log) error {
	tx, err := l.Db.Db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	errIns := l.insertRecord(tx, req)
	if errIns != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
