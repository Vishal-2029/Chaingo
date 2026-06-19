package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Vishal-2029/internal"
	"github.com/Vishal-2029/network"
	"github.com/Vishal-2029/pkg"
)

func main() {
	apiPort := flag.String("api", "8080", "API Port")
	p2pPort := flag.String("p2p", "9000", "P2P Port")
	dbFile := flag.String("db", "chaingo.db", "Database file")
	flag.Parse()

	// Render injects PORT env variable — use it if available
	if envPort := os.Getenv("PORT"); envPort != "" {
		*apiPort = envPort
	}

	db, err := pkg.NewBoltDB(*dbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Printf("Using persistent BoltDB storage: %s\n", *dbFile)

	// Create and set node for networking features
	node := network.NewNode(":" + *p2pPort)
	internal.SetNode(node)

	// NEW: Set database for wallet persistence
	internal.SetDatabase(db)

	go func() {
		node.Start()
	}()

	internal.StartServer(db, *apiPort)
}
