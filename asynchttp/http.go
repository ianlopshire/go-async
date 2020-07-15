package asynchttp

import (
	"net/http"

	"github.com/ianlopshire/go-async"
)

type ResponseFuture struct {
	async.Future

	res *http.Response
	err error
}

func NewResponseFuture() (*ResponseFuture, func(*http.Response, error)) {
	fut, resolve := async.NewFuture()
	resFut := &ResponseFuture{
		Future: fut,
	}

	fn := func(r *http.Response, err error) {
		resFut.res = r
		resFut.err = err
		resolve()
	}

	return resFut, fn
}

func (resFut *ResponseFuture) Result() (*http.Response, error) {
	select {
	case <-resFut.Done():
	default:
		panic("asynchttp: attempting to access un-resolved future")
	}

	return resFut.res, resFut.err
}

type Client struct {
	*http.Client
}

func (c *Client) DoAsync(r *http.Request) *ResponseFuture {
	resFut, resolve := NewResponseFuture()

	go func() {
		resolve(c.Do(r))
	}()

	return resFut
}
