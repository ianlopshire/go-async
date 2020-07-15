# Async

```go
import "github.com/ianlopshire/go-async"
```

Package `async` provides asynchronous primitives and utilities.

## Usage

`Future` is the main primitive provided by `async`.

```go
// Future is a synchronization primitive that guards data that may be available in the
// future.
//
// NewFuture() is the preferred method of creating a new Future.
type Future interface {

	// Done returns a channel that is closed when the Future is resolved. It is safe to
	// call Done multiple times across multiple threads.
	Done() <-chan struct{}
}
```

The provided `Future` primitive does not act as a container for data, but guards data that will be available in the future.

```go
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
```

In practice a `Future` can be embedded into a container to more closely match the traditional behavior of a future.

```go
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

func Example() {
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
```