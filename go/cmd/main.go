package main

import (
	"fmt"
	"log"

	"github.com/leepavelich/projector/go/pkg/cli"
)

func main() {
	opts, err := cli.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	fmt.Printf("opts: %+v\n", opts)
}
