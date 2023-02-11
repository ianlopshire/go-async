package async

import (
	"sync"
)

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

// Latch is a synchronization primitive that can be used to block until a desired state is
// reached.
//
// The zero value for Latch is in an open (blocking) state. Use the package level Resolve
// function to resolve a Latch. Once resolved, the Latch cannot be reopened.
//
// Attempting to resolve a Latch more than once will panic with ErrAlreadyResolved.
//
// A Latch must not be copied after first use.
type Latch struct {
	mu   sync.Mutex
	done chan struct{}
}

// Done returns a channel that will be closed when the Latch is resolved.
func (l *Latch) Done() <-chan struct{} {
	l.mu.Lock()
	if l.done == nil {
		l.done = make(chan struct{})
	}
	d := l.done
	l.mu.Unlock()
	return d
}

func (l *Latch) resolve(fn func()) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if fn != nil {
		fn()
	}

	if l.done == nil {
		l.done = closedchan
		return
	}

	select {
	case <-l.done:
		panic(ErrAlreadyResolved)
	default:
		// Intentionally left blank.
	}

	close(l.done)
}
