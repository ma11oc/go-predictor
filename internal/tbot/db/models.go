package tbot

import (
	v1 "github.com/ma11oc/go-predictor/pkg/api/v1"
	"gorm.io/gorm"
)

// Request is used to store users queries
type Request struct {
	gorm.Model
	ChatID           int64
	MessageID        int
	UserID           int
	Command          string
	CommandArguments string
	Response         *Response
	PersonProfile    *PersonProfile
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

// PersonProfile is used to store request data so it can be reused later
type PersonProfile struct {
	gorm.Model
	RequestID     int
	PersonProfile *v1.PersonProfile `gorm:"embedded"`
}
