package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

const HELLO_RESPONSE = "Hello, from the underworld!\n"

func main() {
	addr := flag.String("addr", ":8081", "address to run the backend server on")
	flag.Parse()

	http.Handle("/hello", handleHello())
	fmt.Printf("starting backend http on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("handled hello request\n")
		io.WriteString(w, HELLO_RESPONSE)
	}
}
