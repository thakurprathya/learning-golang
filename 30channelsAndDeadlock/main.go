package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channelsAndDeadlock in golang")

	// channels are used when we want different goroutines to communitcate (still don't know what's happening inside one another, but can share some signal, etc)

	// myChannel := make(chan int) // creating a channel which will share integer signal
	myChannel := make(chan int, 1) // Buffered Channel  :channel with 1 buffer, that is it allows one extra value to be passed to it without listening

	// myChannel <- 5           // passing value in channel (passing)
	// fmt.Println(<-myChannel) // removing values out of channel (listening)

	// above 3 lines if run as a program will lead to deadlock error
	// channel only allows to pass a value, if something is listening to it (listening first then passing), but reversing also leads to deadlock
	// for program to run properly, goroutines should be handled properly, along with equal number of listners to the values passed

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) { // listening channel (recieve only channel)
		// fmt.Println("Listening to value", <-ch) // if channel has no values will print/output 0, one problem with this is, if we pass 0 signal to channel it will also give same output
		val, isChannelOpen := <-myChannel // to handle above issue, we use this syntax
		fmt.Println("isChannelOpen:", isChannelOpen)
		fmt.Println("val:", val)
		wg.Done()
	}(myChannel, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { // pass in channel (send only channel)
		fmt.Println("putting the values")
		ch <- 0
		ch <- 5
		close(ch) //closing the channel, we cannot pass values to closed channels
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
