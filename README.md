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

## F.A.Q

> Is this package really necessary? The standard library `sync` is already great.

The `sync` package is great! You could use it to solve any problem you could solve with
futures. That being said – in the right circumstances – the futures pattern can lead to
code that some would consider easier to understand and maintain.

At the end of the day, a lot of it comes down to preference. If you and your team don't
like using futures, simply don't use them.

> Why is `Resolve` a function instead of a method of `Future`.

Custom future types embed the `Future` primitive. If `Resolve` was a method of `Future` it
would be exported for every custom future type. Implementing `Resolve` as a package level
function allows implementors of custom future types to choose how to expose resolution
(if at all).
