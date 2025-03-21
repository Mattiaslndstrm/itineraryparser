package main

import "github.com/mattiaslndstrm/iternaryparser/internal/server"

func main() {
	server := server.NewServer(":8080")
	server.Start()
}
