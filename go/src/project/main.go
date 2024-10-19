package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var message string

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/message", MessageHandler).Methods("POST")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GET Request Received")
	fmt.Fprintf(rw, "Hello, %s", message)
}

func MessageHandler(rw http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Message string `json:"message"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(rw, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	message = requestBody.Message
	log.Println("Message updated to:", message)

	fmt.Fprintf(rw, "Message updated to: %s", message)
}
