package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice as the res body
// param w: provides methods for assembling an HTTP response and sending it to a user
// param r: pointer to a struct that holds info about the current request
func home(w http.ResponseWriter, r *http.Request) {
	// since servemux treats the "/" route as a catch-all, we can do this to restrict the route to just "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, err := w.Write([]byte("Hello"))
	if err != nil {
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// extract snippet id from url query string if it exists `/snippet/view?id=1`
	// convert it to an int and make sure it's at least 1
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	_, err = fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	if err != nil {
		return
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// only handle create (POST)
	if r.Method != http.MethodPost {
		// give user info about what request methods are available
		w.Header().Set("Allow", http.MethodPost)

		//// can only call w.WriteHeader() once per response
		//w.WriteHeader(405)
		//
		//// if no WriteHeader call beforehand, w.Write() will automatically send a 200 OK status code
		//_, _ = w.Write([]byte("Method Not Allowed"))

		// or we can just use this helpful function (calls above methods under the hood):
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	_, err := w.Write([]byte("Create a new snippet..."))
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)                    // with servemux, "/" is a catch-all - acts like a wild card, i.e. /**
	mux.HandleFunc("/snippet/view", snippetView) // fixed path, will only strictly match
	mux.HandleFunc("/snippet/create", snippetCreate)

	// if you wanted to, you could use:
	// http.HandleFunc("/", home)
	// which uses `DefaultServeMux` behind the scenes, which is:
	// `var DefaultServeMux = NewServeMux()`
	// avoid this, since it's basically a global var that could be comprised by a third-party package
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
