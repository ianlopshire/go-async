package async_test

import (
	"testing"

	"github.com/ianlopshire/go-async"
)

func TestLatch_alreadyResolved(t *testing.T) {
	// When an already-resolved Latch is resolved it should panic with ErrAlreadyResolved.
	defer func() {
		if recover() != async.ErrAlreadyResolved {
			t.Fatal("expected latch to panic with ErrAlreadyResolved")
		}
	}()

	var counter int
	inc := func() { counter++ }

	l := new(async.Latch)
	async.Resolve(l, inc)
	async.Resolve(l, inc)

	if counter > 1 {
		t.Fatal("expect resolve func to be called at most once")
	}
}
