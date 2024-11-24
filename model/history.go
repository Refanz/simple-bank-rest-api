package model

import "time"

type History struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}
