package main

import (
	"fmt"
	"os"
	
	"github.com/Tech-Preta/aws-resources/pkg/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
