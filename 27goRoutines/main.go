package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// better version of time.sleep
// it has 3 methods:
// ADD (everytime we create goroutine, add it to it, and it will makesure code doesn't end before all the methods are executed)
// DONE (to mark method as done) & WAIT (for waiting the methods)
var waitGroup sync.WaitGroup

func main() {
	fmt.Println("goRoutines in golang")

	// goroutine is a lightweight thread managed by the Go runtime
	// Threads are managed by OS, while goroutines are managed by go runtime
	// Threads have fixed stack of 1mb size, goroutines have flexible stack of 2kb

	// concurency : it is about managing and organizing tasks, one task at a time, but we switch context between them
	// parallelism : it is about executing tasks simultaneously to improve performance

	// real life example : get involved in some task that we forgot about the other in parallelism, to see other in effect we can utilise sleep method
	// without any sleep method or some other wait like waiting for api result, what is happening is, main method is killing its execution before executing go thread
	// go print("Hello")
	// print("World")

	websites := []string{
		"https://google.com",
		"https://youtube.com",
		"https://reqres.in/api/users",
		"https://github.com",
		"https://fb.com",
	}

	// concurently getting statusCode for each site
	// for _, site := range websites {
	// 	getStatusCode1(site)
	// }

	// parallely getting statusCode for each site, by running go routine for each site
	for _, site := range websites {
		go getStatusCode2(site)
		waitGroup.Add(1)
	}

	waitGroup.Wait() // will not allow main to end
}

func print(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(s, i+1)
	}
}

func getStatusCode1(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d code for %s\n", res.StatusCode, endpoint)
}

func getStatusCode2(endpoint string) {
	defer waitGroup.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d code for %s\n", res.StatusCode, endpoint)
}
