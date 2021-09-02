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
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello friend\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func ContextWithKey(ctx context.Context, s string){
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found", s)
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

	ctx := context.WithValue(context.Background(), "language", "go")
	ContextWithKey(ctx, "language")

	fmt.Println("Access and wait: http://localhost:8090/hello")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

