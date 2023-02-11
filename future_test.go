package async_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ianlopshire/go-async"
)

func TestFuture_alreadyResolved(t *testing.T) {
	// When an already-resolved Future is resolved it should panic with ErrAlreadyResolved.
	defer func() {
		if recover() != async.ErrAlreadyResolved {
			t.Fatal("expected Future to panic with ErrAlreadyResolved")
		}
	}()

	fut := new(async.Future[string])
	async.ResolveFuture(fut, "Hello World!", nil)
	async.ResolveFuture(fut, "Hello World!", nil)
}

func TestFuture_Value(t *testing.T) {
	for name, tt := range map[string]struct {
		v   string
		err error
	}{
		"with value": {"Hello World!", nil},
		"with error": {"", errors.New("error")},
	} {
		t.Run(name, func(t *testing.T) {
			fut := new(async.Future[string])
			async.ResolveFuture(fut, tt.v, tt.err)

			v, err := fut.Value()
			if err != tt.err {
				t.Fatalf("Value() unexpected error have %v, want %v", err, tt.err)
			}
			if v != tt.v {
				t.Fatalf("Value() unexpected value have %v, want %v", v, tt.v)
			}
		})
	}
}

func TestFuture_ValueCtx(t *testing.T) {
	for name, tt := range map[string]struct {
		v   string
		err error
	}{
		"with value": {"Hello World!", nil},
		"with error": {"", errors.New("error")},
	} {
		t.Run(name, func(t *testing.T) {
			fut := new(async.Future[string])
			async.ResolveFuture(fut, tt.v, tt.err)

			v, err := fut.ValueCtx(context.Background())
			if err != tt.err {
				t.Fatalf("Value() unexpected error have %v, want %v", err, tt.err)
			}
			if v != tt.v {
				t.Fatalf("Value() unexpected value have %v, want %v", v, tt.v)
			}
		})
	}

	t.Run("with canceled context", func(t *testing.T) {
		timeout := time.After(time.Second)
		done := make(chan bool)

		go func() {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			fut := new(async.Future[string])
			_, err := fut.ValueCtx(ctx)
			if err != context.Canceled {
				t.Errorf("ValueCtx() unexpected error have %v, want %v", err, context.Canceled)
			}

			done <- true
		}()

		select {
		case <-timeout:
			t.Fatal("ValueCtx() future should not have blocked")
		case <-done:
		}
	})
}
