package main

import (
	"fmt"
	"log"
	"net/http"
	"google.golang.org/appengine" // Required external App Engine library
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
	fmt.Fprintln(w, "Hello, from Tenx Infotech")
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
	// Starts the server to receive requests
	appengine.Main()
}
