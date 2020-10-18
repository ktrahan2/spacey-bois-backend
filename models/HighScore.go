package models

//HighScore data structure
type HighScore struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}
