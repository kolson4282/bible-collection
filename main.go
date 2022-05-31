package main

import (
	"os"

	"github.com/kolson4282/bible-collection/memcollection"
	"github.com/kolson4282/bible-collection/server"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	server.StartServer(memcollection.NewMemoryCollection(), port)
}
