package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/portfolio/server"
)

func main() {
	fmt.Println("Starting server...")

	err := server.CreateServer()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	} else {
		fmt.Println("Server started on http://localhost:8080")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}
