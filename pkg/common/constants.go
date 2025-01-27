package common

import "time"

type MessageRequest struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}
