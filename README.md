# Async

```go
import "github.com/ianlopshire/go-async"
```

Package `async` provides asynchronous primitives and utilities.

## Usage

`Latch` is a synchronization primitive that can be used to block until a desired state is reached.

The zero value for `Latch` is in an open (blocking) state. Use the package level `Resolve` function to resolve a Latch. Once resolved, the Latch cannot be reopened.


#### Custom Future 

`Latch` is useful for implementing Futures.

```go
type User struct {
	ID   int
	Name string
}

type UserFuture struct {
	async.Latch
	User User
	Err  error
}

func ResolveUser(fut *UserFuture, user User, err error) {
	async.Resolve(fut, func() {
		fut.User, fut.Err = user, err
	})
}

// This example demonstrates how Latch can be embedded into a custom type to create a
// Future implementation.
func ExampleLatch_customType() {
	userFut := new(UserFuture)

	// Simulate long computation or IO by sleeping before and resolving the future.
	go func() {
		time.Sleep(500 * time.Millisecond)
		user := User{ID: 1, Name: "John Does"}

		ResolveUser(userFut, user, nil)
	}()

	// Block until the future is resolved.
	async.Await(userFut)

	fmt.Println(userFut.User, userFut.Err)
	// output: {1 John Does} <nil>
}
```

## F.A.Q

> Is this package really necessary? The standard library `sync` is already great.

The `sync` package is great! You could use it to solve any problem you could solve with
futures. That being said – in the right circumstances – the futures pattern can lead to
code that some would consider easier to understand and maintain.

At the end of the day, a lot of it comes down to preference. If you and your team don't
like using futures, simply don't use them.

> Why is `Resolve` a function instead of a method of `Latch`.

Custom future types embed the `Latch` primitive. If `Resolve` was a method of `Latch` it
would be exported for every custom future type. Implementing `Resolve` as a package level
function allows implementors of custom future types to choose how to expose resolution
(if at all).
