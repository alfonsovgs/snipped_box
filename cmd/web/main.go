package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Define a new command-line flag the name 'addr', a default value of ":4000"
	// amd some short help text explain what the flag controls. The value of the
	// flag will be strored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Important, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assings it to the addr
	// variable. You need to call this before you see the add variable
	// otherwise it will contain the default value of ":40000". If any errors are
	// encounter during parsing the application will be terminated.
	flag.Parse()

	// Use the http.NewServerMux() function to initialize a new servermux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server wich servers files out of the "./ui/static" directory.
	// Note that the path given to the htt.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching pathsm, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servermux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
