package async_test

import (
	"fmt"
	"time"

	"github.com/ianlopshire/go-async"
)

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
