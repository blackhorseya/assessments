package worker

import (
	"glasnostic/internal/app/entities"
)

// ReadyNotifier declare ready notify
type ReadyNotifier interface {
	WorkerReady(chan []string)
}

// Worker define a worker struct
type Worker struct {
	ID    int
	In    <-chan []string
	Out   chan<- *entities.Profile
	Stop  chan<- bool
	Ready ReadyNotifier
}

// NewWorker new a worker
func NewWorker(ID int, in chan []string, out chan<- *entities.Profile, stop chan<- bool, ready ReadyNotifier) *Worker {
	worker := &Worker{ID: ID, In: in, Out: out, Stop: stop, Ready: ready}

	go func() {
		for {
			ready.WorkerReady(in)
			select {
			case line := <-in:
				p, err := entities.NewProfileFromLine(line)
				if err != nil {
					continue
				}
				out <- p
			}
		}
	}()

	return worker
}
