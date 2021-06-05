package counter

import (
	"errors"
	"sync/atomic"

	"github.com/blackhorseya/ip-rate-limit/pkg/contextx"
	"golang.org/x/time/rate"
)

var (
	ips  = make(map[string]*int32)
	zero = int32(0)
)

type impl struct {
}

// NewImpl serve caller to create an IBiz
func NewImpl() IBiz {
	return &impl{}
}

func (i *impl) Count(ctx contextx.Contextx, ip string, limiter *rate.Limiter) (int, error) {
	count, ok := ips[ip]
	if !ok {
		ips[ip] = &zero
		count = ips[ip]
	}

	if !limiter.Allow() {
		atomic.StoreInt32(count, 0)
		return 0, errors.New("too many request")
	}

	ret := atomic.AddInt32(count, 1)

	return int(ret), nil
}
