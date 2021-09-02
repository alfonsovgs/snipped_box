package main

import (
	"flag"
	"log"
	"net/http"
	"os"
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

	// Use log.New() to create a logger for writting information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writting error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevan
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servermux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil
	infoLog.Printf("Starting server on %s", *addr)

	// Call the ListenAndServe() method on our new http.Server struct.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
