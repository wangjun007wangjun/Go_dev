package goroute

func Add(a, b int, c chan int) {
	sum := a + b
	c <- sum
}
