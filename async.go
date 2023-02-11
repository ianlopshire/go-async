// Package async provides asynchronous primitives and utilities.
package async

import (
	"context"
	"errors"
)

// ErrAlreadyResolved indicates that a Resolver has already been resolved.
var ErrAlreadyResolved = errors.New("async: already resolved")

// Awaiter is an interface that can await a resolution.
type Awaiter interface {
	Done() <-chan struct{}
}

// Await blocks until r is resolved.
//
// synonymous with:
// 	<-a.Done()
func Await(a Awaiter) {
	<-a.Done()
}

// AwaitCtx blocks until a is resolved or the context is canceled.
func AwaitCtx(ctx context.Context, a Awaiter) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-a.Done():
		return nil
	}
}

// Resolver is an interface that wraps the unexported resolve method.
//
// It is implemented by Latch and can also be implemented by embedding Latch into a custom
// type.
type Resolver interface {
	resolve(func())
}

// Resolve resolves a Resolver and runs the given function. It is guaranteed that fn will
// be call at most once.
//
// Calling Resolve more than once for a single Resolver will panic with ErrAlreadyResolved.
func Resolve(r Resolver, fn func()) {
	r.resolve(fn)
}
