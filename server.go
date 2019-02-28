package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func startAsyncServer(resp chan<- string, wg *sync.WaitGroup, port int) *http.Server {
	srv := &http.Server{Addr: ":" + strconv.Itoa(port)}

	go func() {
		defer wg.Done()
		fmt.Printf("Server listening on port %v\n", strconv.Itoa(port))
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server ended in error: %s", err)
		}

		resp <- "Server"
	}()

	return srv
}
