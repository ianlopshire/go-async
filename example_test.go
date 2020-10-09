package async_test

import (
	"errors"
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

func ExampleFuture_select() {
	fut := new(async.ValueFuture)

	// The channel returned by Done() can be used directly in a select statement.
	select {
	case <-fut.Done():
		fmt.Println(fut.Value, fut.Err)
	default:
		fmt.Println("Future not yet resolved")
	}

	// Simulate long computation or IO by sleeping before setting the value and resolving
	// the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		async.ResolveValueFuture(fut, "Hello World!", nil)
	}()

	// Block until the future is resolved.
	async.Await(fut)

	fmt.Println(fut.Value, fut.Err)
	// output:
	// Future not yet resolved
	// Hello World! <nil>
}

func ExampleErrFuture() {
	fut := new(async.ErrFuture)

	// Simulate long computation or IO by sleeping before setting the value and resolving
	// the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		async.ResolveErrFuture(fut, errors.New("there was an error"))
	}()

	// Block until the future is resolved.
	async.Await(fut)

	fmt.Println(fut.Err)
	// output: there was an error
}

func ExampleValueFuture() {
	fut := new(async.ValueFuture)

	// Simulate long computation or IO by sleeping before setting the value and resolving
	// the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		async.ResolveValueFuture(fut, "Hello World!", nil)
	}()

	// Block until the future is resolved.
	async.Await(fut)

	fmt.Println(fut.Value, fut.Err)
	// output: Hello World! <nil>
}
