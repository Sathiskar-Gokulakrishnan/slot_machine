package models

import "time"

type Game struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	PlayerID  string    `json:"player_id" bson:"player_id"`
	WinAmount int       `json:"win_amount" bson:"win_amount"`
	Outcome   string    `json:"outcome" bson:"outcome"`
	PlayedAt  time.Time `json:"played_at" bson:"played_at"`
}
