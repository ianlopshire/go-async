// Package async provides asynchronous primitives and utilities.
package async

// Await blocks until a future resolves.
func Await(f Future) {
	<-f.Done()
	return
}
