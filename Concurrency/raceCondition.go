package Concurrency

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func mainRaceCondtion() {

	var x = 0

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go increment(&x)
		wg1.Add(1)
		go decrement(&x)

	}
	wg1.Wait()
	fmt.Println(x)

	c := make(chan int)
	go prod(1, 2, c)
	go prod(3, 4, c)
	a := <-c
	b := <-c
	fmt.Println(a * b)
}

func increment(x *int) {
	defer wg1.Done()
	fmt.Println("Increasing s")

	*x = *x + 1
}

func decrement(x *int) {
	fmt.Println("Decreasing s")

	*x = *x - 1
	wg1.Done()
}

func prod(v1, v2 int, c chan int) {
	c <- v1 * v2
}

/*
	Race codition is a bug in a program where two or more process try to access and modify a shared
	part of the program (variable ) and when they are not able to do it in a proper manner and fail
	to get the proper result.

	In the above program the race codition occurs between the goroutines running increment and decremenet.
	If you run the proram multiple time, you will get different results because when goroutine1 tries to
	change the value of x, the goroutine2 may have been locking the resource and gorouter1 gets the old
	values of x.


*/
