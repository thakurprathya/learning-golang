package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("raceCondition in golang")

	// Race Condition: Memory Space is a single thing, and consider there are large number of goroutines, if we try to write in the memory using one thread, if another thread wants to write in memory it will cause issue
	// -> go run --race .  :command is used to analysis if there is race condition in program or not, if it exists code exit with status 66
	// mutex can be used to handle race condition

	waitGroup := &sync.WaitGroup{} // pointer
	mutex := &sync.RWMutex{}       // pointer, read-write mutex

	var score = []int{0}

	waitGroup.Add(4)
	// immediately invoked functions (iif)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("1st routine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(waitGroup, mutex)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("2nd routine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(waitGroup, mutex)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("3rd routine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(waitGroup, mutex)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("4th routine")
		m.RLock()
		fmt.Println("score within func:", score)
		m.RUnlock()
		wg.Done()
	}(waitGroup, mutex)

	waitGroup.Wait()

	mutex.RLock()
	fmt.Println("score:", score)
	mutex.RUnlock()
}
