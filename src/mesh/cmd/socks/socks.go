package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/skycoin/skycoin/src/mesh/app"
	"github.com/skycoin/skycoin/src/mesh/messages"
	network "github.com/skycoin/skycoin/src/mesh/nodemanager"
)

func main() {

	//messages.SetDebugLogLevel()
	messages.SetInfoLogLevel()

	var (
		err error
	)

	args := os.Args
	if len(args) < 1 {
		//		printHelp()
		return
	}

	hopsStr := os.Args[1]

	if hopsStr == "--help" {
		//		printHelp()
		return
	}

	hops, err := strconv.Atoi(hopsStr)
	if err != nil {
		fmt.Println("\nThe first argument should be a number of hops\n")
		return
	}

	if hops < 1 {
		fmt.Println("\nThe number of hops should be a positive number > 0\n")
		return
	}

	meshnet := network.NewNetwork()
	defer meshnet.Shutdown()

	clientAddr, serverAddr := meshnet.CreateSequenceOfNodes(hops + 1)

	server, err := app.NewSocksServer(meshnet, serverAddr, "127.0.0.1:8001")
	if err != nil {
		panic(err)
	}

	client, err := app.NewSocksClient(meshnet, clientAddr, "0.0.0.0:8000")
	if err != nil {
		panic(err)
	}

	err = client.Dial(serverAddr)
	if err != nil {
		panic(err)
	}

	server.Serve()

}
