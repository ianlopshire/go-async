package asynchttp_test

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ianlopshire/go-async"
	"github.com/ianlopshire/go-async/asynchttp"
)

func ExampleClient_DoAsync() {
	client := &asynchttp.Client{
		Client: http.DefaultClient,
	}

	r, err := http.NewRequest("GET", "https://proxy.golang.org/github.com/ianlopshire/go-async/@v/v0.1.0.info", nil)
	if err != nil {
		log.Fatal(err)
	}

	resFut := client.DoAsync(r)

	async.Await(resFut)

	res, err := resFut.Result()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// output: {"Version":"v0.1.0","Time":"2020-07-15T18:41:55Z"}
}
