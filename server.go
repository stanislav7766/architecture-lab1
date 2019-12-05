package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"log"
)

func main() {
		http.HandleFunc("/", homeHandler)
		http.HandleFunc("/time", timeHandler)
		defer http.ListenAndServe(":8795", nil)
		fmt.Println("Server is listening on 8795 port...")
}

func timeHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json") 
	jsonTime, err := json.Marshal(map[string]string{"time": time.Now().Format(time.RFC3339)})
	// Handle error if it exists
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// Sending time
	if _, err := res.Write(jsonTime); err != nil {
		log.Printf("Error writing response to the client: %s", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
}
