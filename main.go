// harbormaster/main.go
// Write semantic Dockerfiles, ship performant ones.
// Usage:
//

package main

import (
	"fmt"
	"os"
)

func main() {
	// Parse command line args for a path to a Dockerfile
	if len(os.Args) == 1 {
		fmt.Println("Missing argument: path to Dockerfile")
		os.Exit(1)
	}

	dockerfilePath := os.Args[1]

	fmt.Println("Dockerfile path:", dockerfilePath)
}
