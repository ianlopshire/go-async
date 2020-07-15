package async

import (
	"sync"
)

// Future is a synchronization primitive that guards data that may be available in the
// future.
//
// NewFuture() is the preferred method of creating a new Future.
type Future interface {

	// Done returns a channel that is closed when the Future is resolved. It is safe to
	// call Done multiple times across multiple threads.
	Done() <-chan struct{}
}

// ResolveFunc resolves a future.
type ResolveFunc func()

type future struct {
	mu   sync.Mutex
	done chan struct{}
}

// NewFuture returns a new future and function that will resolve it.
//
// Calling resolve more than once will cause a panic.
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
		panic("async: future is already resolved")
	default:
	}

	close(f.done)
}

func (f *future) Done() <-chan struct{} {
	f.mu.Lock()
	if f.done == nil {
		f.done = make(chan struct{})
	}
	d := f.done
	f.mu.Unlock()
	return d
}
