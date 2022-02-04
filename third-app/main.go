package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}

	rStr := "Hello %s from third app!\n---\n\n"

	// Call other service
	otherServiceURL := os.Getenv("OTHER_SERVICE_URL")
	if otherServiceURL != "" {
		b, err := get(otherServiceURL)
		rStr = rStr + "ERROR: " + err.Error() + "\n"
		rStr = rStr + "RESPONSE: " + b + "\n"
	}

	fmt.Fprintf(w, rStr, name)
}
