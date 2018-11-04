package main

import "fmt"

func list(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}
func main() {
	list(10)
}
