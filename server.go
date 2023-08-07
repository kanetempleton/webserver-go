
package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	// Define a handler for the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	// Define a handler for the "/hello" URL
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, WebGo!")
	})

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
