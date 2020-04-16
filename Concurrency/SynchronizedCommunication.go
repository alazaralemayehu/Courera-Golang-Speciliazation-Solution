package Concurrency

import "fmt"

func main12() {
	c := make(chan int, 2)
	var x = 1

	for {
		go decrement2(x, c)

		go increment2(x, c)
	}
	for i := range c {
		fmt.Println(i)
	}
}

func increment2(x int, c chan int) {
	x = x + 1
	c <- x
}

func decrement2(x int, c chan int) {
	x = x - 1
	c <- x
}
