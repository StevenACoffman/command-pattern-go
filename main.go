package main

import (
	"fmt"
	"os"

	"github.com/StevenACoffman/command-pattern-go/cmd"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
	// exitSuccess is the exit code if the program succeeds.
	exitSuccess = 0
)

func main() {
	// pass all arguments without the executable name
	if err := cmd.Run(os.Args[1:]); err != nil {

		fmt.Printf("Got error when ending %+v\n\n", err)
		os.Exit(exitFail)

	}
	fmt.Println("Successful completion")
	os.Exit(exitSuccess)
}
