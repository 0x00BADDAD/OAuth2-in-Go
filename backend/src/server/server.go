package server

// Aim is to find the way to write an http response to be an html file

import (
	"fmt"
	"log"
	"net/http"
)

func initServer() {
	h1 := func(w http.ResponseWriter, req *http.Request) {
		// Access the headers received
		headersMap := req.Header
		fmt.Println("Printing the header Map: %v\n", headersMap)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
