package models

type KafkaOrder struct {
	Number string `json:"number"`
	UserID string `json:"user_id"`
	Status string `json:"status"`
}