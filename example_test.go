package async_test

import (
	"fmt"
	"time"

	"github.com/ianlopshire/go-async"
)

func ExampleLatch() {
	l := new(async.Latch)

	// Simulate long computation or IO by sleeping before setting the value and resolving
	// the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		async.Resolve(l, nil)
	}()

	// Block until the Latch is resolved.
	async.Await(l)
	fmt.Println("Done!")

	// output: Done!
}

func ExampleLatch_select() {
	l := new(async.Latch)

	// The channel returned by Done() can be used directly in a select statement.
	select {
	case <-l.Done():
		fmt.Println("Done!")
	default:
		fmt.Println("Latch not yet resolved")
	}

	// output: Latch not yet resolved
}
