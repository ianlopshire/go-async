# Async

```go
import "github.com/ianlopshire/go-async"
```

Package `async` provides asynchronous primitives and utilities.

## Usage

`Future` is the main primitive provided by `async`.

```go
type Future interface {
	Done() <-chan struct{}
}
```

The provided `Future` primitive does not act as a container for data, but guards 
data that will be available in the future.

```go
// Define a place where data will be stored.
var (
	value string
	err   error
)

// Create a new Future.
future, resolve := async.NewFuture()

go func() {
	// Sleep to simulate a cpu intensive task or network request.
	time.Sleep(500 * time.Millisecond)
	value = "Hello World!"

	// Resolve the future now that the data has been stored.
	resolve()
}()

// Wait for the future to resolve.
<-future.Done()

// Now that the future is resolved it is safe to use the data.
if err != nil {
log.Fatal(err)
}
fmt.Println(value)
```

In practice a `Future` can be embedded into a type to more closely match the traditional
behavior of a future.

```go
type Container struct {
	Future
	Value string
	Err   error
}

func NewContainer() (*Container, func(string}, error)) {
	f, r := NewFuture()

	v := &Container{
		Future: f,
	}

	return v, func(value string , err error) {
		v.Value = value
		v.Err = err
		r()
	}
}

func DoSomething() {
	// Create a container with an embedded Future.
	container, resolve := NewContainer()

	go func() {
		// Sleep to simulate a cpu intensive task or network request.
		time.Sleep(500 * time.Millisecond)

		// Resolve the container with some data.
		resolve("Hello World!", nil)
	}()

	// Wait for the future to resolve.
	Await(container)

	// Now that the future is resolved it is safe to use the data.
	if err := container.Err; err != nil {
		log.Fatal(err)
	}
	fmt.Println(container.Err)
}
```