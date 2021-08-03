package tbot

import "gorm.io/gorm"

// Request is used to store users queries
type Request struct {
	gorm.Model
	ChatID           int64
	MessageID        int
	UserID           int
	Command          string
	CommandArguments string
	Response         *Response
}

// Response is used to store response payload
type Response struct {
	gorm.Model
	RequestID int
	MessageID int
	ChatID    int64
	TTL       int
	Payload   []byte
}

// User is used to store users info
type User struct {
	gorm.Model
	ID           int `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	UserName     string
	LanguageCode string
	IsBot        bool
}
