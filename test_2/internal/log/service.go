package log

import (
	"time"

	"github.com/inouttt/stkt/test_2/internal/entity"
)

type Services interface {
	Create(req entity.Log) error
}

type svc struct {
	r Repository
}

func Newservices(r Repository) Services {
	return svc{
		r: r,
	}
}

func (s svc) Create(req entity.Log) error {
	if !req.CreatedAt.Valid {
		req.CreatedAt.Time = time.Now()
	}
	return s.r.Create(req)
}
