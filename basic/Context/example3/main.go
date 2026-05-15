package main

import (
	"context"
	"log"
	"time"
)

func main() {

	//// Simple Context
	//ctx := context.TODO()
	//log.Println(ctx)

	//// Context with Values
	//ctx2 := context.WithValue(ctx, "key1", "value1")
	//log.Println(ctx2)

	// Context with Cancel
	//ctx2, cancelFunc := context.WithCancel(ctx)
	//log.Println(ctx.Err())
	//cancelFunc()
	//log.Println(ctx2.Err() == context.Canceled)

	//deadLine, ok := ctx2.Deadline()
	//log.Println(deadLine)
	//log.Println(ok)
	//t := time.Now().Add(1 * time.Second)
	//context.WithDeadline(ctx, t)
	//log.Println(ctx2.Err())
	//time.Sleep(2 * time.Second)
	//log.Println(ctx2.Err() == context.DeadlineExceeded)

	//// Context with Deadline
	//ctx := context.Background()
	//t := time.Now().Add(1 * time.Second)
	//ctx2, _ := context.WithDeadline(ctx, t)
	//deadLine, ok := ctx2.Deadline()
	//log.Println(ctx2)
	//log.Println(deadLine)
	//log.Println(ok)

	//// Context with TimeOut
	//ctx := context.Background()
	//timeout := 1 * time.Second
	//ctx2, _ := context.WithTimeout(ctx, timeout)
	//log.Println("Context Err:", ctx2.Err())
	//time.Sleep(2 * time.Second)
	//log.Println("Context Err:", ctx2.Err())
	//log.Println(ctx2.Err() == context.DeadlineExceeded)

	// Example of Context Done
	ctx := context.Background()
	timeout := 1 * time.Second
	ctx2, _ := context.WithTimeout(ctx, timeout)
	time.Sleep( 2 * time.Second)
	ctx3 := context.WithValue(ctx2, "whatever", 456)
	fn(ctx2)
	fn(ctx3)
}
func fn(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		log.Print("after 2 sec")
	case <-time.After(1 * time.Second):
		log.Println("after 1 sec")
	case <-ctx.Done():
		log.Print("boom context is done")
	}
}

