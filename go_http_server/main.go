package main

import (
	"fmt"
	"net/http"
	/* Provides HTTP client and server implementations.
	Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:  */)

func handler1(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello Server")
}

func main() {
	http.HandleFunc("/", handler1)
	// Register a handler, or what would be like an endpoint for our server
	http.ListenAndServe(":8000", nil)
	/* Listens on the TCP network address and then calls Serve
	with the controller to handle requests on incoming connections. */
}
