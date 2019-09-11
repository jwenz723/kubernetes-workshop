package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	svcname := "frontend"

	addr := flag.String("addr", ":8080", "address to run the frontend server on")
	backendAddr := flag.String("backend-addr", "http://localhost:8081", "address of the running backend server")
	flag.Parse()

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("handled %s request\n", svcname)
		resp, err := http.Get(fmt.Sprintf("%s/hello", *backendAddr))
		if err != nil {
			http.Error(w, fmt.Errorf("failed to reach backend: %w", err).Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, fmt.Errorf("failed to parse backend response: %w", err).Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Underworld says: %s\n", string(b))
	}

	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("starting %s http on %s\n", svcname, *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
