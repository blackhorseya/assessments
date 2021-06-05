package repo

import (
	"time"

	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	rw *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(rw *sqlx.DB) IRepo {
	return &impl{rw: rw}
}

func (i *impl) Ping(ctx contextx.Contextx) (bool, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	err := i.rw.PingContext(timeout)
	if err != nil {
		return false, err
	}

	return true, nil
}
