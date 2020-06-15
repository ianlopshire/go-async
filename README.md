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