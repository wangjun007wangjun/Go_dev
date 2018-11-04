package main

import (
	"fmt"
	"go_dev/day1/goroute_test/goroute"
)

func main() {
	pipe := make(chan int, 1)
	go goroute.Add(100, 200, pipe)

	sum := <-pipe
	fmt.Println("sum=", sum)
}
