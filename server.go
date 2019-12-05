package main

import (
    "fmt"
		"net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
		defer http.ListenAndServe(":8795", nil)
		fmt.Println("Server is listening on 8795 port...")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello")
}
