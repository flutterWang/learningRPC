package pool

import (
	"errors"
	"io"
	"log"
	"os"
	"sync"
)

// Pool -
type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	signal  chan os.Signal
	close   chan os.Signal
	factory func() (io.Closer, error)
	active  uint
	max     uint
	closed  bool
}

// ErrPoolClosed -
var ErrPoolClosed = errors.New("pool is closed")

// New -
func New(fn func() (io.Closer, error), max uint) (*Pool, error) {
	if max <= 0 {
		return nil, errors.New("size is too small")
	}
	return &Pool{
		factory: fn,
		active:  0,
		max:     max,
		closed:  false,
		res:     make(chan io.Closer, max),
		signal:  make(chan os.Signal, 0),
		close:   make(chan os.Signal, 0),
	}, nil
}

// Get -
func (p *Pool) Get() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("get a resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		if p.active < p.max {
			p.active++
			log.Printf("New resource %d", p.active)
			return p.factory()
		}

		return nil, errors.New("pool is full")
	}
}

func (p *Pool) Signal() {

}

// Close -
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	p.active = 0

	close(p.res)

	for r := range p.res {
		r.Close()
	}
}

// Putback -
func (p *Pool) Putback(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
		log.Println("put it back to pool")
	default:
		log.Println("forget the resource")
		r.Close()
	}
}
