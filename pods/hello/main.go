package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	srv *http.Server
	wg  *sync.WaitGroup
)

func main() {
	wg = new(sync.WaitGroup)
	wg.Add(1)
	http.HandleFunc("/", root)
	http.HandleFunc("/stop", stop)

	//--- server
	srv = &http.Server{Addr: ":10123"}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen error: %s", err.Error())
		}
	}()

	wg.Wait()
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, time now is %v", time.Now())
}

func stop(w http.ResponseWriter, r *http.Request) {
	wg.Done()
	fmt.Fprintf(w, "Server will be stopped %v", time.Now())
}
