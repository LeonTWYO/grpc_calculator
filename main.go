// main.go

package main

import (
	"flag"
	"log"
)

func main() {
	// Parse command-line arguments to determine which component to run
	runServer := flag.Bool("server", false, "Run gRPC server")
	runGateway := flag.Bool("gateway", false, "Run gRPC gateway")
	flag.Parse()

	// Check if only one of the components is specified
	if *runServer && *runGateway {
		log.Fatal("Cannot run both server and gateway simultaneously")
	} else if !(*runServer || *runGateway) {
		log.Fatal("Must specify either -server or -gateway")
	}

	// Run the chosen component
	if *runServer {
		runGRPCServer()
	} else if *runGateway {
		runGRPCGateway()
	}
}
