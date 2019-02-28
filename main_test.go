package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestGameThread(t *testing.T) {
	running := true
	testChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go startGame(testChan, &wg, &running)

	running = false
	wait := <-testChan
	if wait != "Game" {
		t.Errorf("Expected testChan=Game, got " + wait)
	}
}

func TestServerThread(t *testing.T) {
	testChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)

	srv := startAsyncServer(testChan, &wg, 5989)
	srv.Shutdown(context.Background())

	wait := <-testChan
	if wait != "Server" {
		t.Errorf("Expected testChan=Server, got " + wait)
	}
}

func TestServGameWaitGroup(t *testing.T) {
	testChan := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	srv := startAsyncServer(testChan, &wg, 5988)

	running := true
	go startGame(testChan, &wg, &running)

	time.Sleep(1 * time.Second)
	srv.Shutdown(context.Background())
	running = false

	wg.Wait()
}
