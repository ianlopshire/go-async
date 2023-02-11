package async

import (
	"context"
)

// Future is a proxy for a result that is initially unknown.
//
// Use `new(Future[T])` to create a new Future. Use the package level ResolveFuture
// function to resolve a Latch.
//
// A Future must not be copied after first use.
type Future[T any] struct {
	l   Latch
	v   T
	err error
}

// Done returns a channel that will be closed when the Future is resolved.
func (fut *Future[T]) Done() <-chan struct{} {
	return fut.l.Done()
}

// Value blocks until the Future is resolved and returns resulting value and error.
//
// It is safe to call Value multiple times form multiple goroutines.
func (fut *Future[T]) Value() (T, error) {
	Await(fut)
	return fut.v, fut.err
}

// ValueCtx blocks until the Future is resolved or the context is canceled.
//
// If the context is canceled, the returned error will be the context's error. A canceled
// context does not necessarily mean that the Future was not resolved or will not resolve
// in the future.
//
// It is safe to call ValueCtx multiple times form multiple goroutines.
func (fut *Future[T]) ValueCtx(ctx context.Context) (T, error) {
	if err := AwaitCtx(ctx, fut); err != nil {
		var zero T
		return zero, err
	}
	return fut.v, fut.err
}

// ResolveFuture resolves a Future and sets its value and error.
//
// Resolving a Future more than once will panic with ErrAlreadyResolved.
func ResolveFuture[T any](fut *Future[T], v T, err error) {
	Resolve(&fut.l, func() {
		fut.v = v
		fut.err = err
	})
}
