package models

type UserActivity struct {
	UserID    string `json:"user_id"`
	Activity  string `json:"activity"`
	TimeStamp string `json:"timestamp"`
	Action    string `json:"action"`
	Metadata  string `json:"metadata"`
}
