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

	preAlloc bool // 是否在创建 pool 的时候就预创建 workers，默认值：false

	// 当 pool 满的时候，新的 Schedule 调用是否阻塞当前 goroutine。默认值：false
	// 如果 block = false，则 Schedule 返回 ErrNoWorkerAvailInPool
	block bool
}

const (
	defaultCapacity = 16
	maxCapacity     = 64
)

func New(capacity int, opts ...Option) *Pool {
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

	for _, opt := range opts {
		opt(p)
	}

	fmt.Printf("workpool start(preAlloc=%t)\n", p.preAlloc)

	if p.preAlloc {
		for i := 0; i < p.capacity; i++ {
			p.newWorker(i + 1)
			p.active <- struct{}{}
		}
	}

	go p.run()
	return p
}

func (p *Pool) run() {
	idx := len(p.active)

	if !p.preAlloc {
	loop:
		for t := range p.tasks {
			p.returnTask(t)
			select {
			case <-p.quit:
				return
			case p.active <- struct{}{}:
				idx++
				p.newWorker(idx)
			default:
				break loop
			}
		}
	}

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

func (p *Pool) returnTask(t Task) {
	go func() {
		p.tasks <- t
	}()
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
var ErrNoIdleWorkerInPool = errors.New("no idle worker")

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	default:
		if p.block {
			p.tasks <- t
			return nil
		}
		return ErrNoIdleWorkerInPool
	}
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("workerpool freed(preAlloc=%t)\n", p.preAlloc)
}
