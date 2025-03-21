package main

import "github.com/mattiaslndstrm/itineraryparser/internal/server"

func main() {
	server := server.NewServer(":8080")
	server.Start()
}
