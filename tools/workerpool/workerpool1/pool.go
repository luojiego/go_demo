package pool

import (
	"errors"
	"fmt"
	"sync"
)

type Task func()

type Pool struct {
	capacity int

	active chan struct{}
	tasks  chan Task
	wg     sync.WaitGroup
	quit   chan struct{}
}

const (
	defaultCapacity = 16
	maxCapacity     = 64
)

func New(capacity int) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}

	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
	}
	fmt.Printf("workpool start\n")
	go p.run()
	return p
}

func (p *Pool) run() {
	idx := 0

	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			// create a new worker
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) newWorker(idx int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker [%03d]: recover panic[%s] and exit\n", idx, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker [%03d]: start\n", idx)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker [%03d]: exit\n", idx)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker [%03d]: receive a task\n", idx)
				t()
				fmt.Printf("worker [%03d]: Finished\n", idx)
			}
		}
	}()
}

var ErrWorkerPoolFreed = errors.New("workerpool freed")

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) Free() {
}
