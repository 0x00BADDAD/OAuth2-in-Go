package main

// Aim is to find the way to write an http response to be an html file

import (
    "net/http"
    "fmt"
    "log"
)

func main() {
    h1 := func(w http.ResponseWriter, req *http.Request) {
        // Access the headers received
        headersMap := req.Header
        fmt.Println("Printing the header Map: %v\n", headersMap)
    }
    http.HandleFunc("/", h1);
    log.Fatal(http.ListenAndServe(":8080", nil));


}
