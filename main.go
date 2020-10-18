package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllHighScores(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit")
	json.NewEncoder(w).Encode(HighScores)
}

func handleRequest() {
	http.HandleFunc("/highscores", returnAllHighScores)
	http.HandleFunc("/", homePage)
	username, exists := os.LookupEnv("USERNAME")
	if exists {
		log.Println(username)
	}
	log.Fatal(http.ListenAndServe(":7000", nil))
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	HighScores = []highscore{
		highscore{ID: 1, Username: "ktrain", Score: 100},
		highscore{ID: 2, Username: "nnpalmore", Score: 200},
	}
	handleRequest()
}

//HighScores is an array for all highscores
var HighScores []highscore

type highscore struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}
