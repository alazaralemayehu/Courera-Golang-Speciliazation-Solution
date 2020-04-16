package Concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex } // creating ChopSticks
type philosopher struct {       //Here are the Philospher which has nunmber
	name      int
	leftChop  *ChopS
	rightChop *ChopS
}

var wg sync.WaitGroup

//This functions simulate the eat functionality of the philospher

func (p philosopher) eat(c chan int) {
	//when the program starts executing, it will put a value into a channel which has two places
	c <- 1
	for i := 0; i < 3; i++ {
		// the following two lines of code generates random numbers
		rand.Seed(time.Now().UnixNano())
		randNumber := rand.Intn(2)

		fmt.Println("Started eating ", p.name)
		//if the generated number is 1 then, the left Chopstick is lifted first
		//if the generated number is not, the right chopstick is lifted first
		if randNumber == 1 {
			p.leftChop.Lock()
			p.rightChop.Lock()
		} else {
			p.rightChop.Lock()
			p.leftChop.Lock()
		}

		//Putting the Chopstics Down
		p.rightChop.Unlock()
		p.leftChop.Unlock()

		fmt.Println("Finished eating ", p.name)

	}

	//This line of Code writes out of the channel Meaning the other third philosopher will start to eat
	<-c
	//Decreasing the number of waiting group
	wg.Done()

}

func main() {
	//Make channel with 2 storage (Buffered Channel)
	channel := make(chan int, 2)
	//Make the waitgroup 5 which represent the 5 philosophers
	wg.Add(5)
	//Creating Chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	//Creating Chopstic Instance
	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{i, CSticks[i], CSticks[(i+1)%5]}
	}

	// this will call the eat methods of all the philosopher with the Buffered Channel
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(channel)
	}

	wg.Wait()

}
