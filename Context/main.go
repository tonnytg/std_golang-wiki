package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func ContextWithTime() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <- time.After(1 * time.Second):
		fmt.Println("time is over")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	fmt.Println("Context Examples")

	ContextWithTime()

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

