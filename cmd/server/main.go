package main

import "github.com/mattiaslndstrm/itineraryparser/internal/server"

const Port = "8080"

func main() {
	server := server.NewServer(":" + Port)
	server.Start()
}
