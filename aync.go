// Package async provides asynchronous primitives and utilities.
package async

import (
	"sync"
)

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

// Awaiter defines the methods that must be implemented to wait on a future.
//
// This interface exists so that the Await function can act on custom future types. See
// ErrFuture and ValueFuture as an Example.
type Awaiter interface {
	Done() <-chan struct{}
}

func Await(f Awaiter) {
	<-f.Done()
}

// Resolver defines the methods that must be implemented to resolve a future.
//
// This interface exists so that the Resolve function can act on custom future types. See
// ErrFuture and ValueFuture as an Example.
type Resolver interface {
	resolve(func())
}

// Resolve resolves a Future and runs fn while the Future is locked.
//
// Calling Resolve more than once for a single future will cause a panic.
//
// This allows a closure to be created that resolves custom Future types. See
// ResolveErrFuture and ResolveValueFuture for an example.
func Resolve(fut Resolver, fn func()) {
	fut.resolve(fn)
}

// Future is a primitive intended to be embedded in custom future types. See ErrFuture and
// ValueFuture for an example.
type Future struct {
	mu   sync.Mutex
	done chan struct{}
}

// resolve resolves f and runs fn while f is locked.
func (f *Future) resolve(fn func()) {
	f.mu.Lock()
	defer f.mu.Unlock()

	fn()

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

// Done returns a channel that is closed when the Future is resolved. It is safe to
// call Done multiple times across multiple threads.
func (f *Future) Done() <-chan struct{} {
	f.mu.Lock()
	if f.done == nil {
		f.done = make(chan struct{})
	}
	d := f.done
	f.mu.Unlock()
	return d
}

// ErrFuture is a future that holds an error.
type ErrFuture struct {
	Future
	Err error
}

// ResolveErrFuture resolves an ErrFuture with the provided error.
func ResolveErrFuture(f *ErrFuture, err error) {
	Resolve(f, func() {
		f.Err = err
	})
}

// ValueFuture is a future that holds an interface{} value and an error.
type ValueFuture struct {
	Future
	Value interface{}
	Err   error
}

// ResolveValueFuture resolves a ValueFuture with the provided value and error.
func ResolveValueFuture(f *ValueFuture, value interface{}, err error) {
	Resolve(f, func() {
		f.Value = value
		f.Err = err
	})
}
