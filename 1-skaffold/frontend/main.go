package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const HELLO_RESPONSE = "Underworld says: %s\n"

func main() {
	addr := flag.String("addr", ":8080", "address to run the frontend server on")
	backendAddr := flag.String("backend-addr", "http://localhost:8081", "address of the running backend server")
	flag.Parse()

	http.Handle("/hello", handleHello(*backendAddr))
	fmt.Printf("starting frontend http on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleHello(backendAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("handled request\n")
		resp, err := getBackendData(backendAddr)
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
		fmt.Fprintf(w, HELLO_RESPONSE, string(b))
	}
}

func getBackendData(backendAddr string) (*http.Response, error) {
	return http.Get(fmt.Sprintf("%s/hello", backendAddr))
}
