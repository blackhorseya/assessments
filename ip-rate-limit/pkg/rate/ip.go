package rate

import (
	"sync"

	"golang.org/x/time/rate"
)

// IP is a ip rate limiter
type IP struct {
	ips   map[string]*rate.Limiter
	mu    *sync.Mutex
	limit rate.Limit
	burst int
}

// NewIP serve caller to new an IP rate limiter
func NewIP(limit rate.Limit, burst int) *IP {
	return &IP{
		ips:   make(map[string]*rate.Limiter),
		mu:    &sync.Mutex{},
		limit: limit,
		burst: burst,
	}
}

// Limiter serve caller get limiter by IP
func (i *IP) Limiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, ok := i.ips[ip]
	if !ok {
		i.mu.Unlock()
		return i.addIP(ip)
	}

	i.mu.Unlock()
	return limiter
}

func (i *IP) addIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.limit, i.burst)
	i.ips[ip] = limiter

	return limiter
}
