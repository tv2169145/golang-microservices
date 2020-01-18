package main

import (
	"fmt"
	"time"
)

func main() {
	now1 := time.Now().String()
	now2 := time.Now().UTC().String()
	zone, _ := time.LoadLocation("Asia/Taipei")
	now3 := time.Now().In(zone).Add(2 * time.Hour).Unix()
	fmt.Println(now1)
	fmt.Println(now2)
	fmt.Println(now3)
}

func helloWorld(done chan struct{}) {
	defer close(done)
	fmt.Println("Hello World!")
}

func foo(done chan struct{}) {
	defer close(done)
	fmt.Println("foo")
}


