package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	// Static files route

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	// we neuter the file system with some middleware to avoid users getting access to directory file listings
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Application routes

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// custom type that embeds http.FileSystem
type neuteredFileSystem struct {
	fs http.FileSystem
}

// Open implement open method on custom fs type
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	// f.Stat() returns FileInfo
	s, err := f.Stat()
	if s.IsDir() {
		// if dir, check for index.html file and open it
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			// Close on the original file to avoid a file descriptor leak
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			// err will be returned and transformed into a 404 by http.FileServer
			return nil, err
		}
	}

	// return the file and let http.FileServer do its thing
	return f, nil
}
