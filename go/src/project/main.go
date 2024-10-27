package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages/{id}", DeleteMessage).Methods("DELETE")
	router.HandleFunc("/api/messages/{id}", UpdateMessage).Methods("PATCH")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CreateMessage(rw http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Message string `json:"message"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(rw, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	message := Message{Text: requestBody.Message}
	if err := DB.Create(&message).Error; err != nil {
		http.Error(rw, "Failed to save message", http.StatusInternalServerError)
		return
	}

	log.Println("Message saved to DB:", message.Text)
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(message)

}

func GetMessage(rw http.ResponseWriter, r *http.Request) {
	var messages []Message

	if err := DB.Find(&messages).Error; err != nil {
		http.Error(rw, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}

	log.Printf("All messages retrieved: %v\n", messages)

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(messages)
}

func DeleteMessage(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := DB.Unscoped().Where("id = ?", id).Delete(&Message{}).Error; err != nil {
		http.Error(rw, "Failed to delete message", http.StatusInternalServerError)
		return
	}

	log.Printf("Message with ID %s deleted", id)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Message with ID %s deleted", id)
}

func UpdateMessage(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var requestBody struct {
		Message string `json:"message"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(rw, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := DB.Model(&Message{}).Where("id = ?", id).Update("text", requestBody.Message).Error; err != nil {
		http.Error(rw, "Failed to update message", http.StatusInternalServerError)
		return
	}

	log.Printf("Message with ID %s updated to: %s", id, requestBody.Message)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Message with ID %s updated to: %s", id, requestBody.Message)
}
