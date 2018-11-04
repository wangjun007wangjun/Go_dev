package main

import (
	"fmt"
	"go_dev/day1/package_example/calc"
)

func main() {
	sum := calc.Add(1, 2)
	sub := calc.Sub(1, 2)

	fmt.Println("Sum=", sum, "Sub=", sub)
}
