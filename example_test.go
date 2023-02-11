package async_test

import (
	"fmt"
	"log"
	"time"

	"github.com/ianlopshire/go-async"
)

func ExampleFuture() {
	// Create a new Future
	fut := new(async.Future[string])

	// Simulate long computation or IO by sleeping before setting the value and resolving
	// the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		async.ResolveFuture(fut, "Hello World!", nil)
	}()

	// Block until the Future is resolved.
	v, err := fut.Value()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v, err)
	// output: Hello World! <nil>
}

func ExampleFuture_select() {
	fut := new(async.Future[string])

	// The channel returned by Done() can be used directly in a select statement.
	select {
	case <-fut.Done():
		fmt.Println(fut.Value())
	default:
		fmt.Println("Future not yet resolved")
	}

	// output: Future not yet resolved
}

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
