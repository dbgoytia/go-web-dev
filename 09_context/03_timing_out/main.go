package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	res, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Println(w, res)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		// Ridiculous long running task
		uid := ctx.Value("userID").(int)
		time.Sleep(4 * time.Second) // Force cancel of context

		// Check context to avoid leakage
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	// If the context got cancelled, then this is going to hit True.
	case <-ctx.Done():
		return 0, ctx.Err()
	// Otherwise it will pull out the value from the channel, but we're not going to get here.
	case i := <-ch:
		return i, nil
	}

}
