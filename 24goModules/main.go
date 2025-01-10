package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules in golang")

	// Here will learn about go mod command which we're extensively using till now
	// -> go mod init github.com/thakurprathya/gomodules  : complete & correct intitialization command
	// first function of go.mod file is it defines what version of go is used in the code

	// it also defines all the other dependencies used in the code, suppose we want to inlcude some library in our code we can do it using go get command
	// -> go get -u github.com/gorilla/mux  : command for mux library
	// it utilised require directive to define external libraries, with direct & indirect check flags (indirect - not currently used, and direct currently used)
	// it also creates go.sum file, which is used to map the hash and checks for the repositories we're pulling data from (security file)

	// we cannot directly view the libraries we pulled, to view them:
	// 1. run -> go env
	// 2. copy GOPATH
	// 3. open the directory
	// 4. migrate further to ./pkg/mod/cache  : location where all the libraries from the web are stored, whenever we use them they are pulled from here
	// 5. migrate further to cache/download : here we can find downloaded libraries
	// 6. now migrating to download/github.com : as above gorilla is fetched from github

	greet()

	// handling requests
	r := mux.NewRouter()                        // Creates a new router instance using the gorilla/mux package, router is responsible for HTTP request routing in your web application
	r.HandleFunc("/", serveHome).Methods("GET") // only serving get method

	// setting up web server
	log.Fatal(http.ListenAndServe(":3000", r)) // will log anything if something goes wrong

	// after running the file we can check route / on localhost:3000
	// also we noticed even after server running we didn't see direct in go.mod
	// all the go mod operations are expensive
	// -> go mod tidy :command (suggested by go), used for directly using the package, removing all unnecessary packages, etc (expensive command)

	// verifying the modules
	// -> go mod verify (returns verify string)

	// listing down the modules
	// -> go list (returns base level modules)
	// -> go list all (returns everything that is installed)
	// -> go list -m all (all used packages are returned)
	// -> go list -m -versions github.com/gorilla/mux (returns all the available versions of a package)

	// -> go mod why github.com/gorilla/mux (asking why we're dependent on this package)
	// -> go mod graph (expensive operation) pulls whole dependency graph
	// -> go mod edit (used to edit go.mod file, if don't wanna touch go.mod manually)
	// -> go mod edit -go 1.16 : set's go version to 1.16
	// -> go mod edit -module : changing module

	// for production builds, where we are making changes in libraries and then pushing them go provides go mod vendor (till now everything is pulled from cache when required)
	// -> go mod vendor (creates a vendor directory (simlar to node_modules)) [used for packaging everything]

	// -> go run main.go -> pull everything from web or cache only not from vendor
	// -> go run -mod=vendor main.go -> first checks vendor for files, then other places

	// https://go.dev/ref/mod
}

func greet() {
	fmt.Println("Morning")
}

func serveHome(w http.ResponseWriter, r *http.Request) { // common syntax while designing backend
	// all the request parameters, urls, etc can be find inside r
	// w is for writing (sending the result)
	w.Write([]byte("<h1>Welcome to go world</h1>"))
}
