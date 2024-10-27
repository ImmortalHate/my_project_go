package main

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Text string `json:"text"` // Наш сервер будет ожидать json c полем text
}
