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
