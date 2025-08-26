package main

import (
	"os"
	"os/signal"

	"github.com/portfolio/server"
)

func main() {
	server.CreateServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}
