package main

import (
	"log"
	"github.com/chiragsoni81245/codeforces-rss/internal/server"
)

func main() {
	log.Println("Starting RSS server on port 8080...")
	server.StartServer()
}

