package main

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine" // Required external App Engine library
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
	}

	fmt.Fprintln(w, "Hello, From Tenx Infotech!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
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
