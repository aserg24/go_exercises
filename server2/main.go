// Server2 is a minimal "echo" and counter server.
package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	s := http.Server{Addr: "localhost:8000"}
	fmt.Println("listening")
	go s.ListenAndServe()
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	s.Shutdown(context.Background())
	go func() {
		for range time.Tick(time.Second) {
			fmt.Println("server stopped")
		}
	}()
	buf.Scan()
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
