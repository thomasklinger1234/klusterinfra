package main

import (
	"log"
	"net/http"
)

func main() {
	httpEndpoint := "0.0.0.0:8080"

	http.HandleFunc("/-/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UP"))
	})

	go func() {
		log.Printf("server listening on http://%s", httpEndpoint)
		err := http.ListenAndServe(httpEndpoint, http.DefaultServeMux)
		if err != nil {
			log.Printf("failed to listen on http://%s: %s", err)
		}
	}()

	select {}
}
