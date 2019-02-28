package main

import (
	"context"
	"flag"
	"sync"
)

func main() {
	port := flag.Int("port", 5454, "Set the port of your server")
	flag.Parse()

	var waitSync sync.WaitGroup
	waitSync.Add(2)
	shutdownSource := make(chan string, 2)
	server := startAsyncServer(shutdownSource, &waitSync, *port)

	gameRunning := true
	go startGame(shutdownSource, &waitSync, &gameRunning)

	endedBy := <-shutdownSource

	if endedBy == "Game" {
		println("Exit signal recieved from game\nShutting down server")
		if err := server.Shutdown(context.Background()); err != nil {
			panic(err)
		}
		end := <-shutdownSource
		println(end + " ended properly")
	} else {
		gameRunning = false
		println("Exit signal recieved from server\nShutting down game")
		end := <-shutdownSource
		println(end + " ended properly")
	}

	waitSync.Wait()
	close(shutdownSource)
	println("Ending program")
}
