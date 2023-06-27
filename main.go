package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice as the res body
// param w: provides methods for assembling an HTTP response and sending it to a user
// param r: pointer to a struct that holds info about the current request
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	// initialize a new servemux
	// register the home function as a handler for "/" url pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// start a new web server
	// two params:
	// 1.TCP network address to listen on
	// 2. servemux created above
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
