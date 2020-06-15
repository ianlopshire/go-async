package async

import (
	"sync"
)

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

// Future guards data that will be available at some point in the future.
type Future interface {
	Done() <-chan struct{}
}

// ResolveFunc resolves a future.
type ResolveFunc func()

type future struct {
	mu   sync.Mutex
	done chan struct{}
}

// NewFuture returns a new future and function that will resolve it.
func NewFuture() (Future, ResolveFunc) {
	f := new(future)
	return f, f.resolve
}

func (f *future) resolve() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.done == nil {
		f.done = closedchan
		return
	}

	select {
	case <-f.done:
		panic("future is already resolved")
	default:
	}

	close(f.done)
}

// Done returns a channel that's closed once the future is resolved.
func (f *future) Done() <-chan struct{} {
	f.mu.Lock()
	if f.done == nil {
		f.done = make(chan struct{})
	}
	d := f.done
	f.mu.Unlock()
	return d
}
