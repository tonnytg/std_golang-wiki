package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converge)

	//for v := range converge {
	//	fmt.Println("valor recebido:", v)
	//}

	for {
		select {
		case x, ok := <-converge:
			if ok {
				fmt.Printf("Value %d was read.\n", x)
			} else {
				fmt.Println("Channel closed!")
			}
		default:
			//fmt.Println("No value to read, exit.")
			// When goroutine finish reading channel, select set to default,
			// then stay waiting come another value to start work
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func envia(p, i chan int) {
	x := 10
	for j := 0; j < 3; j ++ {
		for n := 0; n < x; n++ {
			if n%2 == 0 {
				p <- n
			} else {
				i <- n
			}
		}
		time.Sleep( 5 * time.Second)
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {

	go func() {
		wg.Add(1)
		for v := range p {
			c <- v * 2
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for v := range i {
			c <- v * 3
		}
		wg.Done()
	}()

}

