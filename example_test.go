package async_test

import (
	"fmt"
	"time"

	"github.com/ianlopshire/go-async"
)

func ExampleFuture() {
	// Create a new Future
	fut := new(async.Future)

	// Simulate long computation or IO by sleeping before setting the value and resolving
	// the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		async.Resolve(fut, nil)
	}()

	// Block until the future is resolved.
	async.Await(fut)

	fmt.Println("Done")
	// output: Done
}
