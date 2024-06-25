package models

type Player struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Name    string `json:"name"`
	Credits int    `json:"credits"`
	Status  string `json:"status"`
}
