package async_test

import (
	"fmt"
	"time"

	"github.com/ianlopshire/go-async"
)

type Container struct {
	async.Future
	Value interface{}
	Err   error
}

func NewContainer() (*Container, func(interface{}, error)) {
	fut, resolve := async.NewFuture()
	container := &Container{
		Future: fut,
	}

	fn := func(value interface{}, err error) {
		container.Value = value
		container.Err = err
		resolve()
	}

	return container, fn
}

// This example demonstrates how a Future can be embedded into a container.
func ExampleFuture_container() {
	// Define a value that will be available in the future.
	container, resolve := NewContainer()

	// Simulate long computation or IO by sleeping before and resolving the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		resolve("Hello World!", nil)
	}()

	// Block until the future is resolved.
	async.Await(container)

	fmt.Println(container.Value, container.Err)
	// output: Hello World! <nil>
}
