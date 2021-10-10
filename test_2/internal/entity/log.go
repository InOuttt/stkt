package entity

import (
	"database/sql"
)

const (
	LogTableName = "log"
)

type Log struct {
	Id           string       `field:"id"`
	RequestBody  string       `field:"request_body"`
	ResponseBody string       `field:"response_body"`
	CreatedAt    sql.NullTime `field:"created_at"`
}

func (lr Log) ToInsertQry() []interface{} {
	var qry []interface{}
	qry = append(qry, lr.RequestBody)
	qry = append(qry, lr.ResponseBody)
	qry = append(qry, lr.CreatedAt.Time.Format("2006-01-02 15:04:05"))
	return qry
}
