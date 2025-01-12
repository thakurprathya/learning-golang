package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var waitGroup sync.WaitGroup // here a variable, ideally pointer is created
var mutex sync.Mutex         // here a variable, ideally pointer is created

func main() {
	fmt.Println("mutex in golang")

	// mutex is a mutual execution lock
	// when we need to lock some memory space for some goroutine we utilizes mutex, no other goroutine can use that memory till the first method frees it
	// types mutex, rwmutex (read-write mutex)

	websites := []string{
		"https://google.com",
		"https://youtube.com",
		"https://reqres.in/api/users",
		"https://github.com",
		"https://fb.com",
	}

	for _, site := range websites {
		go getStatusCode(site)
		waitGroup.Add(1)
	}

	waitGroup.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer waitGroup.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS!! Failed to Handle Endpoint")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint) // taking this as memory unit, which should be accessed exclusively
		mutex.Unlock()

		fmt.Printf("%d code for %s\n", res.StatusCode, endpoint)
	}

}
