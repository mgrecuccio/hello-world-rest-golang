package main

import (
	"fmt"
	"net/http"
	"os" // Required external App Engine library
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello OpenShift for Developers!"
	}
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"google.golang.org/appengine" // Required external App Engine library
// )

// func determineListnerAddress() (string, error) {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		return "", fmt.Errorf("$PORT not set")
// 	}
// 	return ":" + port, nil
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, from Tenx Infotech")
// }

// func main() {
// 	addr, err := determineListnerAddress()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	http.HandleFunc("/", hello)
// 	log.Printf("Listening os %s...\n", addr)

// 	// Starting listening
// 	if err := http.ListenAndServe(addr, nil); err != nil {
// 		panic(err)
// 	}
// 	// Starts the server to receive requests
// 	appengine.Main()
// }
