package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	svcname := "backend"

	addr := flag.String("addr", ":8081", "address to run the backend server on")
	flag.Parse()

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("handled %s request\n", svcname)
		io.WriteString(w, "Hello, from the underworld!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("starting %s http on %s\n", svcname, *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
