# Async

```go
import "github.com/ianlopshire/go-async"
```

Package `async` provides asynchronous primitives and utilities.

## Usage

`Future` is the main primitive provided by `async`.

`Future` easily embeds into custom future types, but `ErrFuture` and `ValueFuture` are included for convenience.

#### Custom Future 

```go
type User struct {
	ID   int
	Name string
}

type UserFuture struct {
	async.Future
	User User
	Err  error
}

func ResolveUserFuture(f *UserFuture, user User, err error) {
	async.Resolve(f, func() {
		f.User, f.Err = user, err
	})
}

// This example demonstrates how a Future can be easily embedded into a custom type.
func ExampleFuture_customType() {
	userFut := new(UserFuture)

	// Simulate long computation or IO by sleeping before and resolving the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		user := User{ID: 1, Name: "John Does"}

		ResolveUserFuture(userFut, user, nil)
	}()

	// Block until the future is resolved.
	async.Await(userFut)

	fmt.Println(userFut.User, userFut.Err)
	// output: {1 John Does} <nil>
}
```

#### Err Future

```go
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
```

#### Value Future

```go
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
```

