/*
The classical Dining philosophers problem, implemented with chopsticks as mutexes.

Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times. The philosophers pick up the chopsticks in any order.In order to eat,
a philosopher must get permission from a host which executes in its own goroutine. The host allows no more than
2 philosophers to eat concurrently. Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
on a line by itself, where <number> is the number of the philosopher.

I've also added messaging for how many times a philosopher has eaten, just for verbosity and clarity.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	number                  int
	leftChopstick, rightChopstick *Chopstick
	eatCount                int
}

type Host struct {
	concurrentEaters int
	sync.Mutex
}

// If a philosopher has not yet eaten thrice, then allow them to eat
func (p *Philosopher) eat(host *Host, wg *sync.WaitGroup) {
	for p.eatCount < 3 {
		host.Mutex.Lock()
		// If there is room for another philosopher to eat
		if host.concurrentEaters < 2 {
			host.concurrentEaters++
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()

			fmt.Printf("starting to eat %d\n", p.number)
			p.eatCount++

			time.Sleep(time.Millisecond * 100) // Simulating eating

			fmt.Printf("finishing eating %d\n", p.number)
			fmt.Printf("Philosopher %d has eaten %d times\n\n", p.number, p.eatCount)
			p.rightChopstick.Unlock()
			p.leftChopstick.Unlock()
			host.concurrentEaters--
			host.Mutex.Unlock()
		// The philosopher has to wait to eat, too many concurrent eaters
		} else {
			host.Mutex.Unlock()
			time.Sleep(time.Millisecond * 100) // Waiting for a chance to eat
		}
	}
	wg.Done()
}

func main() {
	// Create the chopsticks
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// Create the host
	host := Host{concurrentEaters: 0}

	// Create the philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
			eatCount:       0,
		}
	}

	// Create a wait group to synchronize the philosophers
	var wg sync.WaitGroup

	// Start the philosophers
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&host, &wg)
	}

	// Wait for all philosophers to finish eating
	wg.Wait()
}
