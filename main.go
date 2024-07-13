package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/takumi616/ielts-vocabularies-api/infrastructures"
	"github.com/takumi616/ielts-vocabularies-api/infrastructures/database"
)

func test(w http.ResponseWriter, r *http.Request) {
	cfg, _ := infrastructures.NewConfig()

	database.Open(context.Background(), cfg.PgConfig)
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
