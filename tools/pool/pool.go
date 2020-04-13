package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool -
type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

// ErrPoolClosed -
var ErrPoolClosed = errors.New("pool is closed")

// New -
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size is too small")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
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
		log.Println("New resource")
		return p.factory()
	}
}

// Close -
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

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
