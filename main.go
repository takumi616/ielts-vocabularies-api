package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test handler with request path: %s", r.URL.Path)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", test)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("APP_CONTAINER_PORT"),
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
