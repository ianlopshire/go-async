// Package async provides asynchronous primitives and utilities.
package async

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

// Await blocks until a future is resolved.
func Await(f Future) {
	<-f.Done()
	return
}
