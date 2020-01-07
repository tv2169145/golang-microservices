package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <-1
		ch <-2
		ch <-3
		ch <-4
		ch <-5
	}()

	i := 0
	for i < 5 {
		fmt.Println(<-ch)
		i++
	}
}

func helloWorld(done chan struct{}) {
	defer close(done)
	fmt.Println("Hello World!")
}

func foo(done chan struct{}) {
	defer close(done)
	fmt.Println("foo")
}


