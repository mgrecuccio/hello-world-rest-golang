package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func determineListnerAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}

func main() {
	addr, err := determineListnerAddress()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", hello)
	log.Printf("Listening os %s...\n", addr)

	// Starting listening
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
