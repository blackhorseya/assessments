package worker

import (
	"glasnostic/internal/app/entities"
	"time"
)

// Pool declare worker pool
type Pool struct {
	Size    int
	Workers []*Worker

	In   chan []string
	Out  chan *entities.Profile
	Stop chan bool

	WorkerChan chan chan []string

	JobQ    [][]string
	WorkerQ []chan []string
}

// NewPool new a Pool for worker
func NewPool(size int) *Pool {
	return &Pool{
		Size:       size,
		In:         make(chan []string),
		Out:        make(chan *entities.Profile),
		Stop:       make(chan bool),
		WorkerChan: make(chan chan []string),
	}
}

// WorkerReady serve caller to notify a worker is ready
func (p *Pool) WorkerReady(worker chan []string) {
	p.WorkerChan <- worker
}

// Start serve caller to start a pool
func (p *Pool) Start() *Pool {
	go func() {
		for {
			if len(p.JobQ) > 0 && len(p.WorkerQ) == 0 && len(p.Workers) <= p.Size {
				worker := NewWorker(len(p.Workers), make(chan []string), p.Out, make(chan<- bool), p)
				p.Workers = append(p.Workers, worker)
			}

			var (
				activeJob    []string
				activeWorker chan []string
			)

			if len(p.JobQ) > 0 && len(p.WorkerQ) > 0 {
				activeJob = p.JobQ[0]
				activeWorker = p.WorkerQ[0]
			}

			select {
			case job := <-p.In:
				p.JobQ = append(p.JobQ, job)
			case worker := <-p.WorkerChan:
				p.WorkerQ = append(p.WorkerQ, worker)
			case activeWorker <- activeJob:
				p.WorkerQ = p.WorkerQ[1:]
				p.JobQ = p.JobQ[1:]
			case <-time.After(2 * time.Second):
				p.Stop <- true
			}
		}
	}()

	return p
}
