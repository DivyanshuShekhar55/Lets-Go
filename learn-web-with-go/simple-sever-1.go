// following is a simpler version of Go server

package main

import (
	//"fmt"
	"log"
	"net/http"
)

// we can have more data inside the server struct according to our needs
// for complete list visit https://pkg.go.dev/net/http#hdr-Servers

type server struct {
	addr string
}

// following is an interface method on any server
// it contains list of handlers against a list of URLs
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello from server"))

	// in older Go codes you might see switch-cases being used to handle routing
	switch r.URL.Path {
	case "/":
		w.Write([]byte("This is home page path"))

	case "/users":
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("this is GET users api path"))

		case http.MethodPost:
			w.Write([]byte("this is POST users api path"))
		}
	}
}

func first_main() {

	s := &server{addr: ":8080"}

	// the first argument is the address
	// second is the handler to handle the request on that address
	// if we pass the whole server as seond argument then...
	// we need to implement the ServeHTTP interface
	// this interface has all the required handler defined
	log.Fatal(http.ListenAndServe(s.addr, s))

	// from terminal or Postman hit localhost:8080  to see the response
	// also hit different routes defined

}
