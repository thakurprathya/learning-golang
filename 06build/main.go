package main

import "fmt"

func main() {
	fmt.Println("Will learn how to build a project")

	// command to build go projects
	// -> go build
	// above command will create build file

	// if we wanna build a project for some other os we need to specify GOOS in command, (go operating system)
	// -> GOOS="windos" go build
	// above command will build.exe file

	// for mac GOOS="darwin"
	// for linux GOOS="linux"
}
