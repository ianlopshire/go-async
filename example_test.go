package async_test

import (
	"fmt"
	"time"

	"github.com/ianlopshire/go-async"
)

func ExampleFuture() {
	// Define a value that will be available in the future.
	var val string
	fut, resolve := async.NewFuture()

	// Simulate long computation or IO by sleeping before setting the value and resolving
	// the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		val = "Hello World!"
		resolve()
	}()

	// Block until the future is resolved.
	async.Await(fut)

	fmt.Println(val)
	// output: Hello World!
}

func ExampleFuture_select() {
	// Define a value that will be available in the future.
	var val string
	fut, resolve := async.NewFuture()

	// The channel returned by Done() can be used directly in a select statement.
	select {
	case <-fut.Done():
		fmt.Println(val)
	default:
		fmt.Println("Future not yet resolved")
	}

	// Simulate long computation or IO by sleeping before setting the value and resolving
	// the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		val = "Hello World!"
		resolve()
	}()

	// Block until the future is resolved.
	<-fut.Done()

	fmt.Println(val)
	// output:
	// Future not yet resolved
	// Hello World!
}
